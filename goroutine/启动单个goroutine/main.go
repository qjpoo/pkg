package main

import (
	"fmt"
	"sync"
)

// 启动goroutine
var wg sync.WaitGroup  // 优雅的等待,它里面有一个计数器

func main() {
	// 启动单个goroutine
	defer fmt.Println("haha ...")
	wg.Add(2)  // 记录goroutine
	go hello() // 1. 创建了一个goroutine 2. 在新的goroutine中执行函数
	go hello()
	fmt.Println("hello_server, main")
	//time.Sleep(time.Second)
	// 等hello执行完(执行hello函数的那个goroutine执行完)
	wg.Wait() // 阻塞, 一直等 所有的goroutine结束

}

func hello() {
	fmt.Println("hello_server hello_server func ...")
	defer wg.Done() // 计数算减1
}
