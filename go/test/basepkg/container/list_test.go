package basePkgT

import "testing"
import "container/list"

func TestList(t *testing.T) {
	var l = list.New()
	l.Init()
	l.PushFront(1)
	l.PushFront(3)
	t.Log(l)
}
