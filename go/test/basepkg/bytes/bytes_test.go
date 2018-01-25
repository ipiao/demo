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
