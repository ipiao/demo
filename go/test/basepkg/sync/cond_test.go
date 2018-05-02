package basePkgT

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var count = 0

func TestCond(t *testing.T) {
	lock := new(sync.Mutex)
	cond := sync.NewCond(lock)
	var syncChan = make(chan struct{}, 1)
	// 3个消费者
	for i := 0; i < 3; i++ {
		go func(i int) {
			for {
				time.Sleep(time.Microsecond * 50)
				lock.Lock()
				for count == 0 {
					println("waiting")
					cond.Wait()
				}
				count--
				fmt.Printf("消费者%d消费一个，剩余 %d\n", i+1, count)
				// if count == 0 {
				// 	cond.Broadcast()
				// } else {
				cond.Signal()
				// }
				lock.Unlock()
			}
		}(i)
	}

	// 生产者
	for i := 0; i < 3; i++ {
		go func(i int) {
			for {
				time.Sleep(time.Microsecond * 50)
				lock.Lock()
				for count >= 10 {
					println("waiting")
					cond.Wait()
				}
				count++
				fmt.Printf("生产者%d生产了一个，剩余 %d\n", i+1, count)
				// if count == 10 {
				// 	cond.Broadcast()
				// } else {
				cond.Signal()
				// }
				lock.Unlock()
			}
		}(i)
	}
	<-syncChan
}
