package basePkgT

import (
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	r := strings.Title("hello")
	t.Log(r)
}

func TestSplit(t *testing.T) {
	r := strings.Split("", ";")
	t.Logs(r)
}
