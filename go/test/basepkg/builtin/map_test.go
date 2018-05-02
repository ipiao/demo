package basePkgT

import (
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
