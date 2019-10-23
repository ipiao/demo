package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sort"
	"strings"
	"github.com/axgle/mahonia"
)


type Proto struct {
	Seq  uint32
	Ck   uint32
	Len  uint32
	Data []byte
	Ok bool
}

func (p *Proto) String() string {
	return fmt.Sprintf("seq: %d, ch: %d, len: %d, ok: %v", p.Seq, p.Ck, p.Len, p.Ok)
}

func ReadData(conn net.Conn) *Proto {
	p := new(Proto)

	b := make([]byte, 12)
	_, err := io.ReadFull(conn, b)
	if err != nil {
		log.Println("read:", err)
		return nil
	}

	p.Seq = binary.BigEndian.Uint32(b[:4])
	p.Ck = binary.BigEndian.Uint32(b[4:8])
	p.Len = binary.BigEndian.Uint32(b[8:12])

	p.Data = make([]byte, p.Len)
	io.ReadFull(conn, p.Data)

	ckdata:=make([]byte,p.Len)
	copy(ckdata,p.Data)
	switch p.Len%4 {
	case 1:
		ckdata = append(ckdata,0xAB,0xAB,0xAB)
	case 2:
		ckdata = append(ckdata,0xAB,0xAB)
	case 3:
		ckdata = append(ckdata,0xAB)
	}

	sum:=p.Seq
	for i:=0;i<len(ckdata)/4;i++{
		sum =sum^  binary.BigEndian.Uint32(ckdata[i*4:(i+1)*4])
	}
	if sum == p.Ck{
		p.Ok =true
	}
	log.Println(p)
	return p
}

func main() {
	conn, _ := net.Dial("tcp", "challenge.yuansuan.cn:7042")

	bio := bufio.NewReader(conn)

	line, _, _ := bio.ReadLine()
	log.Println(string(line))

	id:=strings.Split(string(line),":")[1]


	conn.Write([]byte( fmt.Sprintf("IAM:%s:%s",id,"530151330@qq.com\n")))
	line, _, _ = bio.ReadLine()
	log.Println(string(line))

	var findata=make([]*Proto,0)
	for  {
		p:=ReadData(conn)
		if p!=nil{
			if p.Ok{
				findata = append(findata,p)
			}
		}else{
			break
		}
	}

	sort.Slice(findata,func(i,j int)bool{
		return findata[i].Seq<findata[j].Seq
	})

	lastData:=make([]byte,0)
	for _,d:=range  findata{
		lastData = append(lastData,d.Data...)
	}

	res:=ConvertToString(string(lastData), "gbk", "utf-8")
	log.Println(res)
}


func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}