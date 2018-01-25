package Test

import (
	"fmt"
	"testing"
)

func TestPassParams(t *testing.T) {
	var m = map[int]int{
		0: 1,
		1: 1,
		2: 2,
	}
	fmt.Println("befor pass:", &m)
	passmap(m)
	fmt.Println("after pass:", &m)

	var s = []int{0, 1, 2, 3}
	fmt.Println("befor pass:", &s)
	passslice(s)
	fmt.Println("after pass:", &s)

}

func passmap(m map[int]int) {
	m[0] = 0
	m[3] = 3
	delete(m, 2)
}

func passslice(s []int) {
	ss := s
	ss[1] = 2
	s[0] = 1
}
