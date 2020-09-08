package main

import "fmt"

// 接收一个ch变量, 类型是chan bool类型的
func recv(ch chan bool)  {
	ret := <- ch  // 阻塞,直到收到ch的信号
	fmt.Println(ret)
}

// main函数也是一个goroutine

func main() {
	// 无缓冲通道

	ch := make(chan bool)  // 创建一个无缓冲的通道,因为没有写容量
	go recv(ch)
	ch <- true
	//fmt.Println("main end ...")
}
