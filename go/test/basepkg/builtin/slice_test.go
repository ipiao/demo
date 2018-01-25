package basePkgT

import (
	"fmt"
	"testing"
)

type Hello struct {
	Name string
}

func TestSlice(t *testing.T) {
	h1 := Hello{
		"tom",
	}
	fmt.Printf("1 %p\n", &h1)
	h2 := Hello{
		"bob",
	}
	fmt.Printf("2 %p\n", &h2)
	hs := []Hello{
		h1, h2,
	}
	fmt.Printf("3.1 %p\n", &hs[0])
	fmt.Printf("3.2 %p\n", &hs[1])
	fmt.Printf("3 %p\n", &hs)

	hs1 := returnslice(hs)
	fmt.Printf("7 %p\n", &(hs1[0]))
	fmt.Printf("8 %p\n", &(hs1[1]))
	fmt.Printf("9 %p\n", &hs1)

}

func returnslice(hs []Hello) []Hello {
	fmt.Printf("4 %p\n", &(hs[0]))
	fmt.Printf("5 %p\n", &(hs[1]))
	fmt.Printf("6 %p\n", &hs)
	hs[1] = Hello{
		"bob1",
	}
	return hs
}
