package basePkgT

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIOCopy(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Fatal(err)
	}
}

func TestIOCopyBuffer(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)
	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		t.Fatal(err)
	}
	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		t.Fatal(err)
	}
}

func TestIOCopyN(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read")
	if _, err := io.CopyN(os.Stdout, r, 5); err != nil {
		t.Fatal(err)
	}
}

func TestIOReadAtLeast(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 33)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", buf)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		t.Log("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		t.Log("error:", err)
	}
}

func TestIOReadFull(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	t.Logf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		t.Log("error:", err)
	}
}
func TestIOReadFullBlock(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	buf := make([]byte, 1024)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	t.Logf("%s\n", buf)
}
