package designPattern

import (
	"fmt"
)

// DecoratorSourceIntf 源接口
type DecoratorSourceIntf interface {
	Method()
}

// DecoratorSource 源
type DecoratorSource struct {
	Name string
}

// Method 方法
func (*DecoratorSource) Method() {
	fmt.Println("method from source")
}

// Decorator 装饰器
type Decorator struct {
	source DecoratorSourceIntf
}

// SetSource 设置源
func (d *Decorator) SetSource(source DecoratorSourceIntf) {
	d.source = source
}

// Method 方法
func (d *Decorator) Method() {
	fmt.Println("before decorator!")
	d.source.Method()
	fmt.Println("after decorator!")
}

// 装饰器对拥有的资源进行装饰
// 代理在进行一定的动作前对自己进行补充
