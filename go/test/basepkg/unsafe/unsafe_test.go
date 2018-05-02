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
