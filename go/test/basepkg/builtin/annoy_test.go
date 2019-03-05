package basePkgT

import (
	"fmt"
	"testing"
	"time"
)

func TestAnnoyF(t *testing.T) {
	type User struct {
		Name string
	}

	fnch := make(chan func(), 10)
	go func() {
		for {
			select {
			case f := <-fnch:
				f()
			}
		}
	}()
	// wait := make(chan struct{})
	u := User{Name: "111"}

	fnch <- func() {
		time.Sleep(time.Millisecond * 100)
		u.Name = "222"
	}

	fnch <- func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(u.Name)
	}

	for {
	}
}
