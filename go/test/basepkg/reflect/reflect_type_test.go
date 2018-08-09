package basePkgT

import (
	"reflect"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestRefType(t *testing.T) {
	type Time time.Time

	now := Time(time.Now())

	v := reflect.ValueOf(now)
	assert.Equal(t, v.Kind(), reflect.Struct)

	assert.Equal(t, v.Type().Name(), "Time")
	assert.Equal(t, v.Type().String(), "basePkgT.Time")

	assert.Equal(t, v.Type().String(), "basePkgT.Time")
}

type User struct {
	Name string
	Age  int
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func TestMtd(t *testing.T) {
	u1 := User{Name: "user1", Age: 20}
	u2 := &User{Name: "user2", Age: 21}

	u1.SetName("u1")
	u2.SetName("u2")
	assert.Equal(t, u1.Name, "user1")
	assert.Equal(t, u2.Name, "u2")
}
