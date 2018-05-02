package basePkgT

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//sum += 1 //1
			atomic.AddUint32(&sum, 1) //2
		}()
	}
	wg.Wait()
	t.Log(sum)
}

func TestCAS(t *testing.T) {
	var syncChan = make(chan bool, 1)
	var value int32 = 2
	go func() {
		for {
			log.Println(11)
			atomic.AddInt32(&value, 2)
			log.Println(12)
		}
	}()

	go func() {
		for {
			log.Println(21)
			v := atomic.LoadInt32(&value)
			if atomic.CompareAndSwapInt32(&value, 20, v+5) {
				syncChan <- true
				break
			}
			log.Println(22)
		}
	}()
	<-syncChan
	t.Log(value)

}
func TestAV(t *testing.T) {
	var cv atomic.Value
	cv.Store([]int{1, 2})
	ans(&cv)
	t.Log(cv.Load())
}

func ans(cv *atomic.Value) {
	cv.Store([]int{2, 3, 4})
}
