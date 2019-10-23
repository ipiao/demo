package basePkgT

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	var strChan = make(chan string, 3)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func() {
		<-syncChan1
		fmt.Println("received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("[receive]", elem)
			} else {
				break
			}
		}
		fmt.Println("Stopped [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[Sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal . [sender]")
			}
		}
		fmt.Println("Wait 2 seconds [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

func TestChan2(t *testing.T) {
	var mapChan = make(chan map[string]int, 1)
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped . [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			// time.Sleep(time.Millisecond)
			fmt.Printf("The count map:%v .[sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func TestChan3(t *testing.T) {
	chanCap := 5
	intChan := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}

func TestChan4(t *testing.T) {
	var syncCh = make(chan struct{})
	var ch1 = make(chan int, 0)
	var ch2 = make(chan int, 0)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
	}()
	go func() {
		for {
			select {
			case <-ch1:
				println(1)
			case <-ch1:
				println(12)
			case <-ch2:
				println(2)
			case <-time.After(time.Second):
				println(3)
				// default:
				// 	println("default")
			}
		}

	}()
	go func() {
		select {
		case <-ch1:
			println(21)
		case <-ch2:
			println(2)
		case <-time.After(time.Second):
			println(3)
			syncCh <- struct{}{}
			// default:
			// 	println("default")
		}
	}()
	<-syncCh
}

func TestChan5(t *testing.T) {
	intChan := make(chan int, 1)
	ticker := time.NewTimer(time.Second)
	go func() {
		for _ = range ticker.C {
			ticker.Reset(time.Second)
			select {
			case intChan <- 1:
				fmt.Printf("1")
			case intChan <- 2:
				fmt.Printf("2")
			case intChan <- 3:
				fmt.Printf("3")
			default:
				fmt.Printf("default")
			}
		}
		fmt.Println("End.[Sender]")
	}()
	var sum int
	for e := range intChan {
		fmt.Println("receive ", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			time.Sleep(2 * time.Second)
			ticker.Stop()

			break
		}
	}
	fmt.Println("End. [rec]")
	time.Sleep(2 * time.Second)

}

func TestChan6(t *testing.T) {
	chCap := 3
	var ch = make(chan int, chCap)
	for i := 0; i < chCap; i++ {
		ch <- i
	}

	select {
	case ch <- chCap:
		t.Log("block")
	default:
		t.Log("default")
	}
	<-ch
}

func TestChan7(t *testing.T) {
	var ch = make(chan int, 1)
	select {
	case ch <- 1:
		t.Log(1)
	default:
		t.Log("default")
	}
	<-ch
}

func TestLock(t *testing.T) {
	lock := sync.RWMutex{}
	lock.RLock()
	lock.RLock()
	lock.RUnlock()
	// lock.Unlock()
}

func TestChanBreak(t *testing.T) {
	for {
		select {
		case c := <-time.After(time.Second):
			log.Println(c)
			continue
		}
	}
}

func TestChanClose(t *testing.T) {
	ch1 := make(chan int, 10)
	for i := 1; i < 20; i++ {
		fmt.Printf("send %d \n", i)
		time.Sleep(time.Millisecond * 100)
		ch1 <- i
	}
	sig := make(chan bool)
	go func() {
		time.AfterFunc(time.Millisecond*500, func() { sig <- true })
	}()
	for {
		select {
		case c := <-time.After(time.Second):
			fmt.Printf("now : %v", c)
		case i, _ := <-ch1:
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		case <-sig:
			close(ch1)
			return
			// return
		}
	}

}

func TestAfter(t *testing.T) {
	ch := make(chan bool, 1)
	// ch <- true
	select {
	case c := <-time.After(0):
		t.Log(c)
	case v := <-ch:
		t.Log(v)

	}
}
