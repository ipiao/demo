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

func TestCopy(t *testing.T) {
	s := make([]int, 50)
	copy(s[3:], s[2:])
	s = s[:len(s)-1]
	t.Log(s)
}

func TestCopy2(t *testing.T) {
	s := make([]int, 50)
	s = append(s[:2], s[3:]...)
	t.Log(s)
}

func TestCap(t *testing.T) {
	s := make([]int, 0, 3)
	s = append(s, 0, 0, 0)
	t.Log(cap(s))
	s = append(s, 1)
	t.Log(cap(s))
}

func TestComp(t *testing.T) {
	s := []int{1, 2, 3, 4}
	s1 := s
	s2 := s
	s2[3] = 5
	t.Log(s1, s2)
}
