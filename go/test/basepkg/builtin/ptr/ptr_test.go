package ptr

import (
	"log"
	"testing"
)

type MyInt int

func (i MyInt) Add(a int) MyInt {
	i = i + MyInt(a)
	return i
}

func (i *MyInt) add(a int) MyInt {
	*i = *i + MyInt(a)
	log.Print(i)
	return *i
}

func TestPtr1(t *testing.T) {
	i := MyInt(3)
	t.Log(&i)
	i1 := i.Add(2)
	t.Log(&i1, i, i1)
}

func TestPtr2(t *testing.T) {
	i := MyInt(3)
	t.Log(&i)
	i1 := i.add(2)
	t.Log(&i1, i, i1)
}
