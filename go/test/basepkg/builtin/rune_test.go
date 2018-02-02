package basePkgT

import "testing"

func TestRune(t *testing.T) {
	s := []string{"中", "国", "人"}
	bs0 := []byte(s[0])
	bs1 := []byte(s[1])
	bs2 := []byte(s[2])
	println(string(append(bs0, append(bs1, bs2...)...)))
}
