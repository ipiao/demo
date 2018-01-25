package deferr

import "testing"

func TestDefer(t *testing.T) {
	var sum int
	for i := 0; i < 5; i++ {
		defer func(a int) {
			t.Log(sum, a)
			sum = sum + a
			t.Log(sum)
		}(sum + i)
	}
}
