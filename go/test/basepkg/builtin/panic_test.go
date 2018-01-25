package basePkgT

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	fmt.Printf("hello world my name is %s, I'm %d\r\n", "songxingzhu", 26)

	go func() {
		myPainc()
	}()
	time.Sleep(1000)
	fmt.Printf("这里应该执行到！")
}
func myPainc() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("出了错：", err)
		}
	}()
	var x = 30
	var y = 0
	//panic("我就是一个大错误！")
	var c = x / y
	log.Println(c)
}
