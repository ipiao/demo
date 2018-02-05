package main

// func main() {
// 	conn1, conn2 := net.Pipe()
// 	defer conn1.Close()
// 	defer conn2.Close()
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			time.Sleep(time.Second * 1)
// 			conn1.Write([]byte("hello"))
// 			conn1.Write([]byte("	"))
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		var buffer bytes.Buffer
// 		go func() {
// 			for {
// 				line, err := buffer.ReadString('\t')
// 				if err != nil {
// 					if err == io.EOF {
// 						continue
// 					}
// 					log.Println("read error ", err)
// 				} else {
// 					log.Printf("read %d bytes from conn1", len(line))
// 				}
// 			}
// 		}()
// 		for {
// 			var bs = make([]byte, 1)
// 			_, err := conn2.Read(bs)
// 			if err != nil {
// 				log.Println("read err ", err)
// 			}
// 			buffer.WriteByte(bs[0])
// 		}
// 	}()
// 	wg.Wait()
// }
