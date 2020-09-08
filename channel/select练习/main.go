package main

import "fmt"

func main() {

	var ch = make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case ret := <-ch:
			fmt.Println(ret)  // 0, 2, 4, 6, 8
		}
	}

}
