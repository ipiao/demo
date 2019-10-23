package basePkgT

import (
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	str := "123456788asdddddddddddddddddddddddddddddddddsfdddsfffffffffffffffffgsdffffffffsfddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"
	size := unsafe.Sizeof(str)
	println(size)
}

func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func TestSliceByteToString(t *testing.T) {
	b := []byte{'a', 'b', 'c', 'd', 'e'}
	s := SliceByteToString(b)
	t.Log(s)
}

func TestStringToSliceByte(t *testing.T) {
	s := "abcde"
	b := StringToSliceByte(s)
	t.Log(b)
}
