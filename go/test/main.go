package main

import (
	"fmt"
	"log"
	"math"
	"net/url"
	"sort"
	"time"

	"github.com/ipiao/metools/mencode"
)

type EmptyInterface interface {
}

type WithFuncInterface interface {
	Func()
}

type TestStruct struct {
	Member int
}

func (t *TestStruct) Func() {
	fmt.Printf("Haha\n")
}
func TestEmptyInterface(i EmptyInterface) {
	fmt.Printf("Interface: %v\n", i == nil)
}

func TestWithFuncInterface(i WithFuncInterface) {
	fmt.Printf("Func Interface: %v\n", i == nil)
}

func TestWithFuncStruct(i *TestStruct) {
	fmt.Printf("Struce Interface: %v\n", i == nil)
}

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}

func rotate(nums []int, k int) {
	temp := make([]int, k, k)
	copy(temp, nums[len(nums)-k:])
	fmt.Println(nums, temp)
	copy(nums[k:], nums[:len(nums)-k])
	fmt.Println(nums, temp)
	copy(nums[:k], temp)
	fmt.Println(nums, temp)

}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var ret = make([]int, 0)
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			ret = append(ret, nums1[i])
			i += 1
			j += 1
		} else {
			if nums1[i] < nums2[j] {
				i++
			} else {
				j++
			}
		}
	}
	return ret
}

func UrlEncode() {

	// str1 := "eyJhbGlfdGlkIjoiMjAxODA3MDkwMDAwMDAwMDEwNTc4MzUwMDAzMDAzMjciLCJsaWQiOiI5NiIsInNpZCI6IjEwMDEyOSIsInRpZCI6Ijc5In0%3D"
	// bs, _ := base64.StdEncoding.DecodeString(str1)
	// var pa = make(map[string]string)
	// json.Unmarshal(bs, &pa)
	// log.Println(string(bs))
	// log.Println(pa)

	str := "eyJhbGlfdGlkIjoiMjAxODA3MDkwMDAwMDAwMDEwNTc4MzUwMDAzMDAzMjciLCJsaWQiOiI5NiIsInNpZCI6IjEwMDEyOSIsInRpZCI6Ijc5In0%3D"
	// var bs = make([]byte, 1024)
	urlll, _ := url.QueryUnescape(str)
	log.Println(urlll)

	// ur3l, _ := url.Parse(urlll)
	// // base64.URLEncoding.Decode(bs, []byte(str))
	// fmt.Println(ur3l.EscapedPath())
}

// func main() {
// 	var test *TestStruct = nil
// 	TestEmptyInterface(test)
// 	TestWithFuncInterface(test)
// 	TestWithFuncStruct(test)
// 	test.Func()
// }

type MyTime time.Time

// func (t MyTime) String() string {
// 	return time.Time(t).Format("2006-01-02 15:04:05")
// }

type TestTime struct {
	Mt time.Time
}

func main() {
	// mt := TestTime{Mt: time.Now()}
	// bs, _ := json.Marshal(mt)
	// fmt.Println(string(bs))

	log.Println(mencode.MD5("gh_27631a767399"))
}
