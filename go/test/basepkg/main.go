package main

import (
	"fmt"
	"hash/crc32"
)

// func main() {
// 	runtime.GOMAXPROCS(1)

// 	var gChan = make(chan string, 20)
// 	go func() {
// 		for s := range gChan {
// 			fmt.Println(s)
// 		}
// 	}()
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer func() {
// 			defer wg.Done()
// 			gChan <- "A done"
// 		}()
// 		for i := 0; i < 3; i++ {
// 			gChan <- fmt.Sprint("A", i)
// 			go func(i int) {
// 				gChan <- fmt.Sprint("AG", i)
// 			}(i)
// 		}
// 	}()
// 	go func() {
// 		defer func() {
// 			defer wg.Done()
// 			gChan <- "B done"
// 		}()

// 		for i := 0; i < 3; i++ {
// 			gChan <- fmt.Sprint("B", i)
// 			go func(i int) {
// 				gChan <- fmt.Sprint("BG", i)
// 			}(i)
// 		}
// 	}()
// 	time.Sleep(time.Second)
// 	wg.Wait()
// 	gChan <- "done waiting"
// 	time.Sleep(time.Second)
// }

type Test struct {
	Name string
}

var list map[string]*Test

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := "ocLDat9znfZRmaJEYidoqWgOlDcM"
	// s := "13083985232"
	// s := "ocLDat8XUjYY17ymP9fVkx74uQVI"
	// s := "ocLDatzbC3uivk6o-dYJK8jb0RS0"
	// s := "oNJFD0czG4zpbRLCH_yE_dvn33qU"
	// s := "ouVCW1Dy-C0rG78SDPyCKUeu761I"
	no := int(crc32.ChecksumIEEE([]byte(s)) % 10)
	fmt.Println(no)
}
