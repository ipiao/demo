package lib

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	"github.com/gorilla/schema"
)

// RolesQuery query for /roles
type RolesQuery struct {
	Limit   int      `schema:"limit" validate:"gt=0"`
	Offset  int      `schema:"offset" validate:"gt=0"`
	Include []string `schema:"include"`
	Filter  Pepple   `schema:"filter"`
}

type Pepple struct {
	Name string
	Age  int
}

type Filter interface{}

func TestSchema(t *testing.T) {
	var val = url.Values{
		"limit":   []string{"1"},
		"offset":  []string{"2"},
		"include": []string{"[\"123\",\"456\"]"},
		"filter": []string{
			"{\"name\":\"tom\",\"age\":1}",
		},
	}
	var query = &RolesQuery{}
	var decoder = schema.NewDecoder()
	decoder.RegisterConverter(Pepple{}, func(value string) reflect.Value {
		f := Pepple{}
		json.Unmarshal([]byte(value), &f)
		return reflect.ValueOf(f)
	})
	err := decoder.Decode(query, val)
	if err != nil {
		panic(err)
	}
	t.Log(query)
}
