package Test

import (
	"encoding/json"
	"testing"
)

func JsonConvert(m interface{}, dest interface{}) (err error) {
	bs, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(bs, dest)
	}
	return
}

func TestInt(t *testing.T) {
	var i interface{} = []int{1, 2}

	bs, _ := json.Marshal(i)

	var ni interface{}

	json.Unmarshal(bs, &ni)

	dest := []int{}

	JsonConvert(ni, &dest)

	t.Log(dest)

	// m, ok := ni.([]interface{})
	// for _, mm := range m {
	// 	ms, ok := mm.(float64)
	// 	t.Log(ni, ms, ok)
	// }
	// t.Log(ni, m, ok)
}
