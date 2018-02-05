package main

import (
	"log"
	"net"
)

// func main() {
// 	parseCIDR()
// }

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
	names, err = net.LookupAddr("192.168.1.168")
	log.Println(names, err)
	names, err = net.LookupAddr("0.0.0.0")
	log.Println(names, err)
	names, err = net.LookupAddr("::1")
	log.Println(names, err)
	names, err = net.LookupAddr("fe80::8d8b:c8c8:2b19:c25b")
	log.Println(names, err)
}

// DNS域
func lookupCNAME() {
	cname, err := net.LookupCNAME("baidu.com")
	log.Println("cname is ", cname, " err is ", err)
}

// 查找主机ip
func lookupHost() {
	addrs, err := net.LookupHost("git.zhixubao.com")
	log.Println(" err is ", err)
	for _, addr := range addrs {
		log.Printf("addr is %s", addr)
	}
}

// 查找ip地址
func lookupIP() {
	ips, err := net.LookupIP("git.zhixubao.com")
	log.Println(" err is ", err)
	for _, ip := range ips {
		log.Printf("ip is %s", ip)
		log.Printf("ip.DefaultMask %s", ip.DefaultMask())
	}
}

// 查找DNS邮件交换记录
func lookupMX() {
	mxs, err := net.LookupMX("zhixubao.com")
	log.Println(" err is ", err)
	for _, mx := range mxs {
		log.Printf("host is %s,pref is %d", mx.Host, mx.Pref)
	}
}

// 名称服务器记录
func lookupNS() {
	nss, err := net.LookupNS("zhixubao.com")
	log.Println(" err is ", err)
	for _, ns := range nss {
		log.Printf("host is %s", ns.Host)
	}
}

// 查找服务端口
func lookupPort() {
	port, err := net.LookupPort("tcp", "7305/chromium-brows")
	log.Println("port is ", port, " err is ", err)
}

func flags() {
	log.Println(net.FlagUp)
}

func parseMac() {
	addr, err := net.ParseMAC("01-23-45-67-89-ab-cd-ef")
	if err != nil {
		log.Println(err)
	}
	log.Println(addr)
}

func parseCIDR() {
	ip, ipnet, err := net.ParseCIDR("127.0.0.1/16")
	log.Println(ip, ipnet, err)
}
