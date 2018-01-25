package basePkgT

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string
	Age  int
	Num  int
}

// Data 运单数据结构
type Data struct {
	Data    interface{} `json:"data,omitempty"`
	Include interface{} `json:"include,omitempty"`
}

func TestJson2(t *testing.T) {
	a := []User{
		User{
			Name: "tom",
		}, User{
			Name: "tony",
		}, User{
			Name: "tommy",
		},
	}
	var data = Data{
		Data: a,
	}
	r, _ := json.Marshal(&data)
	t.Log(string(r))

	var data2 = Data{Data: &[]User{}}

	var s = `{"data":[{"Name":"tom","Age":0,"Num":0},{"Name":"tony","Age":0,"Num":0},{"Name":"tommy","Age":0,"Num":0}]}`
	json.Unmarshal([]byte(s), &data2)
	t.Logf("%+q", data2)
}

func TestJson(t *testing.T) {
	a := []User{
		User{
			Name: "tom",
		}, User{
			Name: "tony",
		}, User{
			Name: "tommy",
		},
	}
	r, _ := json.Marshal(&a)
	t.Log(string(r))

	var as []User
	var s = `[{"name":"tom","age":0,"num":0},{"name":"tony","age":0,"num":0},{"name":"tommy","age":0,"num":0}]`
	json.Unmarshal([]byte(s), &as)
	t.Log(as)
}

func TestJSONDecode(t *testing.T) {
	var u User
	f, _ := os.Open("test2.json")
	dec := json.NewDecoder(f)
	dec.UseNumber()
	dec.Decode(&u)
	assert.Equal(t, u.Name, "tom")
	assert.Equal(t, u.Age, 15)
}

func TestJSONEncode(t *testing.T) {
	var u = User{Name: "jack", Age: 12}
	f, _ := os.Create("test2.json")
	dec := json.NewEncoder(f)
	dec.Encode(&u)
}
