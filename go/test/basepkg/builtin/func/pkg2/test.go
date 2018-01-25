package pkg2

import (
	"log"

	_ "github.com/ipiao/me/test/basepkg/func/pkg"
)

func Hello() {
	log.Println("jhello")
}
