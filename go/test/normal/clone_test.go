package normalTest

import (
	"fmt"
	"testing"

	metools "github.com/ipiao/metools/utils"
)

type UserTestClone struct {
	Name string
	Age  int
	Info *UserInfoTestClone
}

type UserInfoTestClone struct {
	School string
}

// TestClone 测试
func TestClone(t *testing.T) {
	user := &UserTestClone{
		Name: "tom",
		Age:  10,
		Info: &UserInfoTestClone{
			School: "xinglan",
		},
	}
	user2 := *user
	user2.Info.School = "change school"
	fmt.Println("user2 is ", user2)

	fmt.Println("user is ", user)
}

// TestMetoolsClone 测试
func TestMetoolsClone(t *testing.T) {
	user := &UserTestClone{
		Name: "tom",
		Age:  10,
		Info: &UserInfoTestClone{
			School: "xinglan",
		},
	}
	user2 := &UserTestClone{}
	err := metools.DeepCopy(user, user2)
	t.Log(err)
	user2.Info.School = "change school"
	fmt.Println("user2 is ", user2)

	fmt.Println("user is ", user)
}
