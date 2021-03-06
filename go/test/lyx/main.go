package main

import (
	"log"
	"time"

	"github.com/ipiao/remote"
)

func main() {

	done := make(chan struct{})

	fordo()

	<-done
}

func fordo() {
	r := remote.NewRemote("https://nsj-m.yy0578.com")
	ret := make(map[string]interface{})

	for {
		err := r.Post("/v1/bbs/queryDetails", map[string]interface{}{"detailId": "654"}, &ret)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Millisecond * 5)
	}
}

func tickerdo() {
	ticker := time.NewTicker(time.Millisecond * 50)
	for {
		select {
		case <-ticker.C:
			do()
		}
	}

}

func do() {
	r := remote.NewRemote("https://nsj-m.yy0578.com")
	ret := make(map[string]interface{})
	err := r.Post("de", map[string]interface{}{"detailId": "654"}, &ret)
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
}
