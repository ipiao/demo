package basePkgT

import (
	"bufio"
	"bytes"
	"testing"
)

func TestBufferIO(t *testing.T) {
	reader := bytes.NewBufferString("adbndsfnddksfncjxzbxcBDdsfsdfsdfgdsfgsdgdfgewhdfasnfAfnlyfA")
	for i := 0; i < 5; i++ {
		buf := bufio.NewReaderSize(reader, 2)
		line, isprefix, err := buf.ReadLine()
		t.Log(string(line), isprefix, err)
	}
}
