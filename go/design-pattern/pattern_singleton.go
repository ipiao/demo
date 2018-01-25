package designPattern

import (
	"fmt"
	"sync"
)

// Singleton 单例
type Singleton struct {
	name string
}

var (
	singleyonInstance *Singleton
)

// GetSingletonInstance 获取单例
func GetSingletonInstance() *Singleton {
	var lock = &sync.Mutex{}
	if singleyonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleyonInstance = &Singleton{name: "primary name"}
	}
	return singleyonInstance
}

// Hello 单例方法
func (s *Singleton) Hello() {
	fmt.Println("hello from sigleton,my name is ", s.name)
}

// ChangeName 单例方法
func (s *Singleton) ChangeName(name string) {
	var lock = &sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	s.name = name
	fmt.Println("change name is ", name)
}
