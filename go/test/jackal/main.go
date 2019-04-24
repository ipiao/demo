package main

import (
	"crypto/tls"
	"log"

	xmpp "github.com/ipiao/go-xmpp"
)

// xmpp 客户端
type XmppClient struct {
	user, passwd string
	debug        bool
	client       *xmpp.Client
	cmds         chan (*cmd)
}

type cmd struct {
	t int
	d interface{}
}

// 新建一个连接
func (c *XmppClient) renewConn() error {
	host := "127.0.0.1:5222"
	opts := xmpp.Options{
		Host:                         host,
		User:                         c.user,
		Password:                     c.passwd,
		Debug:                        c.debug,
		Session:                      true,
		InsecureAllowUnencryptedAuth: true,
		NoTLS:                        true,
		StartTLS:                     true,
		TLSConfig:                    &tls.Config{InsecureSkipVerify: true},
		NoAuth:                       false,
	}
	client, err := opts.NewClient()
	c.client = client
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// fcm必须要用tls
// user是<sendId>@gcm.googleapis.com
// passwd 就是 appKey
func NewXmppClient(user, passwd string, debug bool) (*XmppClient, error) {
	c := &XmppClient{
		user:   user,
		passwd: passwd,
		debug:  debug,
	}
	err := c.renewConn()
	if err != nil {
		return nil, err
	}
	go c.run()
	return c, nil
}

func (c *XmppClient) run() {
	for {
		chat, err := c.client.Recv()
		if err != nil {
			log.Println(err)
		}
		log.Println(chat)
	}
}

func upload() {

	client, err := NewXmppClient("001@localhost", "1234", true)
	if err != nil {
		log.Fatal("1:", err)
	}
	defer client.client.Close()

	// client.client.

	client.client.SendOrg(`
	<iq from='001@localhost'
		id='step_03'
		to='upload.localhost'
		type='get'>
	<request xmlns='urn:xmpp:http:upload:0'
		filename='test.jpg'
		size='161647'
		content-type='image/jpeg' />
	</iq>`)

	//
	// http.NewRequest("PUT", "/upload", nil)

}

func main() {
	listen()
}

func listen() {

	client, err := NewXmppClient("003@localhost", "1234", true)
	if err != nil {
		log.Fatal("1:", err)
	}
	defer client.client.Close()

	for {
	}
	// client.client.SendOrg(`
	// <iq from='001@localhost'
	// 	id='step_03'
	// 	to='upload.localhost'
	// 	type='get'>
	// <request xmlns='urn:xmpp:http:upload:0'
	// 	filename='test.jpg'
	// 	size='161647'
	// 	content-type='image/jpeg' />
	// </iq>`)

	//
	// http.NewRequest("PUT", "/upload", nil)

}
