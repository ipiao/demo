package main

import (
	"log"
	"time"

	xmpp "github.com/mattn/go-xmpp"
)

func main() {
	// 147.110.83.7
	client, err := xmpp.NewClient("127.0.0.1", "1000000@baozhao.com", "123456", false)
	if err != nil {
		log.Fatal("1:", err)
	}
	defer client.Close()
	// client := xmpp.Client{}
	// err := client.PingC2S("ykk@localhost", "127.0.0.1:5222")
	// if err != nil {
	// 	log.Fatal("1:", err)
	// }
	log.Println("11")
	for {
		time.Sleep(time.Second)
		log.Println("22")
		n, err := client.SendKeepAlive()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n)
	}
}
