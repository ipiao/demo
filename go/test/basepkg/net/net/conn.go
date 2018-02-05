package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func serve() {
	lister, err := net.Listen("udp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	defer lister.Close()
	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Println("Listen Error ", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		str, err := read(conn)
		if err != nil {
			if err == io.EOF {
				log.Println("connetion is closed by annoner side. ")
			} else {
				log.Println("read error ", err)
			}
			break
		}
		resp := fmt.Sprintf("receive msg:%s for client", str)
		n, err := write(conn, resp)
		if err != nil {
			log.Println("write error ", err)
		}
		log.Printf("Sent response (written %d bytes): %s.", n, resp)
	}
}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == '\t' {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte('\t')
	return conn.Write(buffer.Bytes())
}
