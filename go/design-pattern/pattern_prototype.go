package designPattern

import (
	"fmt"
)

// AdapterSource 适配器源
type AdapterSource struct{}

// Method1 源方法
func (AdapterSource) Method1() {
	fmt.Println("this is method1 from source")
}

// AdapterTarget 适配器结果
type AdapterTarget interface {
	Method1()
	Method2()
}

//*****************类适配模式

// Adapter1 适配器实体1 类适配模式
type Adapter1 struct {
	AdapterSource
}

// Method2 方法
func (Adapter1) Method2() {
	fmt.Println("this is method2 from adapter1")
}

//**********对象适配模式

// Adapter2 适配器实体2 对象适配模式
type Adapter2 struct {
	Source AdapterSource
}

// Method1 方法
func (a *Adapter2) Method1() {
	a.Source.Method1()
}

// Method2 方法
func (Adapter2) Method2() {
	fmt.Println("this is method2 from adapter2")
}

//************接口适配模式

// AdapterSource2 适配器源2 接口适配模式
type AdapterSource2 interface {
	Method1()
	Method2()
}

// Adapter31 适配器
type Adapter31 struct {
	AdapterSource2
}

// Method1 方法
func (Adapter31) Method1() {
	fmt.Println("this is method1 from adapter31")
}

// Adapter32 适配器
type Adapter32 struct {
	AdapterSource2
}

// Method2 方法
func (Adapter32) Method2() {
	fmt.Println("this is method2 from adapter32")
}
