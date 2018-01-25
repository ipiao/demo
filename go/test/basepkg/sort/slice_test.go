package basePkgT

import (
	"fmt"
	"sort"
	"testing"
)

type SortTestSlice struct {
	Name string
	Age  int
}

func (p SortTestSlice) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func TestSortSlice(t *testing.T) {
	people := []SortTestBase{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

}
