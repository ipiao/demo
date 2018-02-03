package main

import (
	"log"
	"net"
)

func main() {
	lookupAddr()
}

// 大概是获取系统的网络接口地址
func interfaceAddr() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		log.Printf("Network is %s;String is %s", addr.Network(), addr.String())
	}
}

// 大概是获取系统的网络接口
func interfaces() {
	itfs, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for i, itf := range itfs {
		addrs, _ := itf.Addrs()
		for j, addr := range addrs {
			log.Printf("interface %d addr %d Network is %s;String is %s", i, j, addr.Network(), addr.String())
		}
		log.Printf("Flags %d %s", itf.Flags, itf.Flags.String())
		log.Printf("HardwareAddr %d %s", itf.HardwareAddr, itf.HardwareAddr.String())
		log.Printf("Index %d", itf.Index)
		log.Printf("MTU %d", itf.MTU)
		log.Printf("Name %s", itf.Name)
		addrs2, _ := itf.MulticastAddrs()
		for j, addr := range addrs2 {
			log.Printf("interface %d addrs2 %d MulticastAddrs is %s;String is %s", i, j, addr.Network(), addr.String())
		}
	}
}

// 组装host和port
func joinHostPort() {
	log.Println(net.JoinHostPort("localhost", "5678"))
	log.Println(net.JoinHostPort("ff02::1", "5678"))
}

// 貌似只能识别本地会换的地址
func lookupAddr() {
	names, err := net.LookupAddr("127.0.0.1")
	log.Println(names, err)
	names, err = net.LookupAddr("ff02::1")
	log.Println(names, err)
	names, err = net.LookupAddr("ff01::1")
	log.Println(names, err)
	names, err = net.LookupAddr("ff02::1:ff19:c25b")
	log.Println(names, err)
}
