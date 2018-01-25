package Test

import (
	"fmt"
	"sort"
	"testing"
)

type P struct {
	CloseTime int
	Priority  int
}

func TestSort(t *testing.T) {
	ps := []P{
		{1, 1},
		{2, 1},
		{3, 10},
		{0, 1},
		{0, 10},
		{0, 2},
		{5, 2},
	}

	sort.Slice(ps, func(i, j int) bool {
		//if ps[i].CloseTime > 0 && ps[j].CloseTime > 0 {
		//	if ps[i].CloseTime < ps[j].CloseTime {
		//		return true
		//	}
		//	//if ps[i].CloseTime == ps[j].CloseTime {
		//	//	return ps[i].Priority < ps[j].Priority
		//	//}
		//} else if ps[i].CloseTime == 0 && ps[j].CloseTime == 0 {
		//	return ps[i].Priority < ps[j].Priority
		//} else {
		//	return ps[i].CloseTime > ps[j].CloseTime
		//}
		//return false
		return (-ps[i].CloseTime*100 + 100 - ps[i].Priority) < (-ps[j].CloseTime*100 + 100 - ps[j].Priority)
	})

	fmt.Println(ps)
}
