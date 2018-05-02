package main

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
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D

}
