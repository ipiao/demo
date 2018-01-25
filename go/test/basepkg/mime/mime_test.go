package mime

import (
	"bytes"
	"mime"
	"testing"
)

func TestWord(t *testing.T) {
	enc := new(mime.WordEncoder)
	s := enc.Encode("utf-8", "中国") //=?utf-8??=E4=B8=AD=E5=9B=BD?=
	t.Log(s)

	dec := mime.WordDecoder{}
	reader := bytes.NewBufferString(s)
	dec.CharsetReader("utf-8", reader)
	// o, err := dec.Decode(s)
	// assert.Nil(t, err)
	// t.Log(o)
}
