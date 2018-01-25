package basePkgT

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	a, err := time.ParseInLocation("20060102", "20170102", time.Local)
	t.Log(a, err)
	t.Log(a.Unix())

	ss := strings.Split("2017..2", ".")
	for _, s := range ss {
		t.Log(len(s))
	}
}

func TestParseDuration(t *testing.T) {
	d, _ := time.ParseDuration("7h25m1m")
	t.Log(d, d.Hours())
}

func TestAfter(t *testing.T) {
	//	for {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("timed out")
	}
	//	}
}

func TestTick(t *testing.T) {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v \n", now)
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 1)
	for d := range ticker.C {
		fmt.Println(d)
	}
}

func TestFixedZone(t *testing.T) {
	cn := time.FixedZone("CN", 8)
	t.Log(cn)
	t.Log(time.LoadLocation("Local"))
}

func TestMonth(t *testing.T) {
	t.Log(time.February)
}
