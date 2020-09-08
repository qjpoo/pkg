package main

import "fmt"

func main() {
	// 有缓冲通道的

	ch := make(chan bool, 1)
	ch <- false
	go recv(ch)
	ch <- true
	go recv(ch)
	fmt.Println(len(ch), cap(ch)) // 0, 1  len数据量的大小, cap容量
	fmt.Println("main end ...")

}

func recv(ch chan bool)  {
	ret:= <- ch
	fmt.Println(ret)

}
