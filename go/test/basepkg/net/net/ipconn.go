package main

import (
	"bytes"
	"flag"
	"log"
	"net"
	"time"
)

func main() {
	a := flag.String("p", "ip4:icmp", "protocol")
	b := flag.String("l", "0.0.0.0", "listen ip")
	flag.Parse()
	laddr, err := net.ResolveIPAddr(*a, *b)
	if err != nil {
		log.Fatal(err)

	}
	c, err := net.ListenIP(*a, laddr)
	if err != nil {
		log.Fatal(err)

	}
	for {

		buf := make([]byte, 2048)
		n, addr, err := c.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		ID := int(buf[5]) + int(buf[4])*256
		log.Println(addr, " ", ID, buf[0:n])
	}
}

func ipclient() {
	lip := net.ParseIP("127.0.0.1:7777")
	conn, err := net.DialIP("ip", &net.IPAddr{IP: lip}, &net.IPAddr{IP: lip})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		time.Sleep(time.Second)
		conn.Write([]byte("hello	"))
	}
}

func ipserve() {
	lip := net.ParseIP("127.0.0.1:7777")
	conn, err := net.ListenIP("", &net.IPAddr{IP: lip})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var buffer bytes.Buffer
	for {
		var bs = make([]byte, 1)
		_, ipaddr, err := conn.ReadFromIP(bs)
		if err != nil {
			log.Println("read error ", err)
		}
		log.Println("read ipaddr", ipaddr)
		readByte := bs[0]
		buffer.WriteByte(readByte)
		if readByte == '\t' {
			log.Println(buffer.String())
			buffer.Reset()
		}
	}
}
