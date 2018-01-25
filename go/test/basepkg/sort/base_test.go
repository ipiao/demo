package basePkgT

import (
	"fmt"
	"sort"
	"testing"
)

type SortTestBase struct {
	Name string
	Age  int
}

func (p SortTestBase) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAgSortTestBaseByAgee implements sort.Interface for []Person based on
// the Age field.
// 实现Sort Interface 接口
type SortTestBaseByAge []SortTestBase

func (a SortTestBaseByAge) Len() int           { return len(a) }
func (a SortTestBaseByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortTestBaseByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func TestSortBase(t *testing.T) {
	people := []SortTestBase{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(SortTestBaseByAge(people))
	fmt.Println(people)

}
