package main

import (
	"crypto/tls"
	"log"
	"time"

	xmpp "github.com/ipiao/go-xmpp"
)

func main() {
	opts := xmpp.Options{
		Host:      "127.0.0.1:5222",
		User:      "123@xmpp.ickapay.xx",
		Password:  "123456",
		NoTLS:     true,
		Debug:     false,
		Session:   false,
		StartTLS:  true,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client, err := opts.NewClient()
	if err != nil {
		log.Fatal("1:", err)
	}
	defer client.Close()

	go func() {
		for {
			time.Sleep(time.Second)
			// log.Println("22")
			_, err := client.SendKeepAlive()
			if err != nil {
				log.Fatal(err)
			}
			// log.Println(n)
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			n, _ := client.Send(xmpp.Chat{
				Remote: "ext.xmpp.ykk.xx",
				Type:   "chat",
				Text:   "hello",
				// Subject   string
				// Thread    string
				// Roster    Roster
				// Other     []string
				// OtherElem []XMLElement
				// Stamp     :
			})
			log.Println(n)
		}
	}()

	for {
		chat, err := client.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(chat)
	}

}
