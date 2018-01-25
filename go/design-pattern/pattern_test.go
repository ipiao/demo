package designPattern

import (
	"testing"
)

// // TestFactory 测试简单工厂模式
// func TestFactory(t *testing.T) {
// 	factory := SendFactory{}
// 	// sender := factory.Produce("sms")
// 	sender := factory.ProduceSender("mail")
// 	// sender := factory.Produce("other") will panic
// 	sender.Send()
// }

// // TestFactory 测试多工厂模式
// func TestFactory2(t *testing.T) {
// 	factory := SendFactory{}
// 	sender := factory.ProduceMailSender()
// 	sender.Send()
// }

// // TestAbFactory 抽象工厂
// func TestAbFactory(t *testing.T) {
// 	provide := SmsFactory{}
// 	sender := provide.ProduceSender()
// 	sender.Send()
// }

// // TestSingleton 单例
// func TestSingleton(t *testing.T) {
// 	ins := GetSingletonInstance()
// 	ins.Hello()
// 	ins.ChangeName("a new name")
// 	ins2 := GetSingletonInstance()
// 	ins2.Hello()
// }

// // TestSingleton2 单例并发
// func TestSingleton2(t *testing.T) {
// 	for i := 0; i < 30; i++ {
// 		go func(i int) {
// 			fmt.Println("i is ", i)
// 			ins := GetSingletonInstance()
// 			//ins.Hello()
// 			ins.ChangeName(fmt.Sprintf("a new name %d", i))
// 			ins.Hello()
// 		}(i)
// 	}
// 	time.Sleep(time.Second * 3)
// }

// TestAdapter 适配器
func TestAdapter(t *testing.T) {
	// 类适配
	target1 := Adapter1{}
	target1.Method1()
	target1.Method2()
	// 对象适配
	target2 := Adapter2{}
	target2.Method1()
	target2.Method2()
	// 接口适配
	target31 := Adapter31{}
	target32 := Adapter32{}
	target31.Method1()
	// target31.Method2()
	target32.Method2()
}
