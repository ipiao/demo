package main

import (
	"bytes"
	"log"
	"net"
	"time"
)

func main() {
	go udpserve()
	go udpclient()
	for {
	}
}

func udpserve() {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:12343")
	if err != nil {
		log.Fatal(err)

	}
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)

	}
	defer conn.Close()
	var buffer bytes.Buffer
	for {
		var bs = make([]byte, 1)
		_, ipaddr, err := conn.ReadFromUDP(bs)
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

func udpclient() {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:34234")
	if err != nil {
		log.Fatal(err)

	}
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:12343")
	if err != nil {
		log.Fatal(err)

	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		time.Sleep(time.Second)
		conn.Write([]byte("hello	"))
	}
}
