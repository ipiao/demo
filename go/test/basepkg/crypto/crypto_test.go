package basePkgT

import (
	"crypto"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	t.Log(crypto.MD5.Available())
	t.Log(crypto.MD5.HashFunc())
	t.Log(crypto.MD5.New())
	t.Log(crypto.MD5.Size())
}

func TestMD5(t *testing.T) {
	source := "helloworld"

	h := md5.New()
	h.Write([]byte(source))
	t.Log(hex.EncodeToString(h.Sum(nil)))

	h1 := md5.New()
	h1.Write([]byte(source))
	fmt.Printf("%x", h1.Sum(nil))

}
