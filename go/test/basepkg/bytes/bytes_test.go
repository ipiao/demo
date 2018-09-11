package basePkgT

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	var a = []byte{'a', 'b', 'c', 'a'}
	var b = []byte{'a', 'b', 'c', 'd'}
	t.Log(bytes.Compare(a, b))
	t.Log(bytes.Equal(a, b))
	t.Log(bytes.Contains(a, b))
	t.Log(bytes.ContainsAny(a, "a"))
	t.Log(bytes.ContainsRune(a, 97))
	t.Log(bytes.Count(a, []byte{97}))
	t.Log(bytes.EqualFold([]byte("GO"), []byte("go")))
	t.Logf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
	t.Log(bytes.FieldsFunc([]byte("  foo bar  baz   "), func(r rune) bool {
		return r == ' '
	}))
}

func TestRunes(t *testing.T) {
	s := "中国人"
	b := []byte(s)
	r := bytes.Runes(b)
	for _, bb := range b {
		t.Log(string([]byte{bb}))
	}
	for _, rr := range r {
		t.Log(string([]rune{rr}))
	}
	t.Log(b)
	t.Log(r)
}

// rune  = int32
// byte = uint8
