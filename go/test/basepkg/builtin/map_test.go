package basePkgT

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	ocean := make(map[int]int)
	for i := 0; i < 4; i++ {
		ocean[i] = i
	}
	sea := make(map[int]int)
	for k, v := range ocean {
		sea[k] = v
	}
	sea = ocean
	fmt.Printf("sea %+v,ocean %+v \n", sea, ocean)

	delete(sea, 0)

	fmt.Printf("sea %+v,ocean %+v \n", sea, ocean)
}

type Key struct {
	A int
	B int
}

func (k Key) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d-%d", k.A, k.B)), nil
}

func TestStructKey(t *testing.T) {
	var m = map[Key]int64{
		{A: 1, B: 2}: 12,
		{A: 2, B: 3}: 23,
	}
	bs, err := json.Marshal(m)
	t.Log(m)
	t.Log(err)
	t.Log(string(bs))
}
