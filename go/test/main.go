package main

import (
	"fmt"
	"log"
	"math"
)

type EmptyInterface interface {
}

type WithFuncInterface interface {
	Func()
}

type TestStruct struct {
	Member int
}

func (t *TestStruct) Func() {
	fmt.Printf("Haha\n")
}
func TestEmptyInterface(i EmptyInterface) {
	fmt.Printf("Interface: %v\n", i == nil)
}

func TestWithFuncInterface(i WithFuncInterface) {
	fmt.Printf("Func Interface: %v\n", i == nil)
}

func TestWithFuncStruct(i *TestStruct) {
	fmt.Printf("Struce Interface: %v\n", i == nil)
}

func test() {
	fmt.Printf("hello")
}

// func main() {
// 	var test *TestStruct = nil
// 	TestEmptyInterface(test)
// 	TestWithFuncInterface(test)
// 	TestWithFuncStruct(test)
// 	test.Func()
// }

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}
func main() {
	log.Println("hello")
}
