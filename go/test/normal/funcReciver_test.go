package normalTest

import (
	"fmt"
	"testing"
)

type IntfFunReciver interface {
	SetName()
	SetName2()
}

type UserFunReciver struct {
	Name string
}

func (u *UserFunReciver) SetName() {
	fmt.Printf("SetName接收到的指针：%x", &u)
	u.Name = "name1 recive copy"
}

func (u *UserFunReciver) SetName2() {
	fmt.Printf("SetName2接收到的指针：%x", &*u)
	u.Name = "name2 recive ptr"
}

// TestFunReciveCopy 测试接受器
func TestFunReciveCopy(t *testing.T) {
	u1 := UserFunReciver{"user1"}
	fmt.Printf("u1的指针：%x", &u1)
	fmt.Println("SetName调用前：", u1.Name)
	u1.SetName()
	fmt.Println("SetName2调用后：", u1.Name)

	fmt.Println("SetName2调用后：", u1.Name)
	u1.SetName2() // go auto get ptr
	fmt.Println("SetName2调用后：", u1.Name)
}

// TestFunRecivePtr 测试接受器
func TestFunRecivePtr(t *testing.T) {
	u1 := &UserFunReciver{"user1"}
	fmt.Printf("u1的指针：%x", &*u1)
	fmt.Println("SetName调用前：", u1.Name)
	u1.SetName()
	fmt.Println("SetName2调用后：", u1.Name)

	fmt.Println("SetName2调用后：", u1.Name)
	u1.SetName2()
	fmt.Println("SetName2调用后：", u1.Name)
}

func TestFunReciveCopy2(t *testing.T) {
	var u1 = UserFunReciver{"user1"}
	u1.SetName()
	u1.SetName2()

	// var u2 IntfFunReciver = u1
}
