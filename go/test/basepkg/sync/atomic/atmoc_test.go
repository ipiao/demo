package basePkgT

import (
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
