package basePkgT

import "testing"
import "sync"
import "time"
import "fmt"

var (
	rwlock = sync.RWMutex{}
)

func TestRwlock(t *testing.T) {

	rwlock.RLock()
	fmt.Println(1)
	// go func() {
	rwlock.Lock()
	fmt.Println(3)
	fmt.Println(4)
	rwlock.Unlock()
	// }()
	time.Sleep(time.Second * 1)
	fmt.Println(2)
	rwlock.RUnlock()
}
