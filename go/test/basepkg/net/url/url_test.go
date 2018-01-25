package basePkgT

import (
	"encoding/json"
	"net/url"
	"testing"
)

type User struct {
	Name string
	Age  int
}

type Qyery struct {
	Limit  int
	Offset int
	Filter User
}

func TestValueEncode(t *testing.T) {
	var val = url.Values{
		"limit":  []string{"1"},
		"offset": []string{"2"},
		"filter": []string{"{\"name\":\"tom\",\"age\":20}"},
	}
	t.Log(val.Encode())
}

func TestValueEncode2(t *testing.T) {
	var user = User{
		Name: "tom",
		Age:  20,
	}
	bs, _ := json.Marshal(&user)
	var val = url.Values{
		"limit":  []string{"1"},
		"offset": []string{"2"},
		"filter": []string{string(bs)},
	}
	t.Log(val.Encode())
}

func TestValueParse(t *testing.T) {
	var path = "http://127.0.0.1:1234/roles?filter=%7B%22name%22%3A%22tom%22%2C%22age%22%3A20%7D&limit=1&offset=2"
	val, _ := url.ParseRequestURI(path)
	t.Log(val.Query())
}

func TestValueParse2(t *testing.T) {
	var path = "http://127.0.0.1:1234/roles?filter=%7B%22Name%22%3A%22tom%22%2C%22Age%22%3A20%7D&limit=1&offset=2"
	val, _ := url.ParseRequestURI(path)
	t.Log(val.Query())
}

func TestUrlValueParse(t *testing.T) {
	var val = url.Values{
		"f1": {"f11", "f12"},
	}
	t.Log(val.Encode())
}
