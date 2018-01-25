package basePkgT

import (
	"testing"

	metools "github.com/ipiao/metools/utils"
)

func TestHeap(t *testing.T) {
	var ints = &metools.InterfaceHeap{1, 2, 3, 4}
	ints.Push(5)
	t.Log(ints, len(*ints), cap(*ints))
	ints.Pop()
	t.Log(ints, len(*ints), cap(*ints))
	// heap.Fix(ints, 1)
}
