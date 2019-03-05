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
	host := "47.110.83.7:5222"
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
		NoAuth:                       true,
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
	chat, err := c.client.Recv()
	if err != nil {
		log.Println(err)
	}
	log.Println(chat)
}

func main() {
	// j, _ := jid.New("1000000", "localhost", "xxtest2", true)

	// testStm := stream.NewMockC2S(uuid.New(), j)

	// testStm.SetJID(j)
	// // 147.110.83.7
	// log.Println(testStm)

	client, err := NewXmppClient("123@xmpp.ickapay.com", "1234", true)
	if err != nil {
		// log.Fatal("1:", err)
	}
	defer client.client.Close()

	// client.client.Register(uuid.New().String(), "1000011", "123456")
	// client.client.Send(xmpp.Chat{})
	// for {
	// 	time.Sleep(time.Second)
	// 	log.Println("22")
	// 	n, err := client.SendKeepAlive()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(n)
	// }
}
