package ptr

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

func TestErrNil(t *testing.T) {
	var err = interface{}((*Error)(nil))
	assert.Nil(t, err)
	// t.Log(err)
	he, ok := err.(*Error)
	assert.True(t, ok)
	assert.Nil(t, he)
}
