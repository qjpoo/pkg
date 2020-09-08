package main

import (
	"fmt"
	"runtime"
	"sync"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	defer wg.Done()
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	defer wg.Done()
}

func main() {
	wg.Add(2)
	runtime.GOMAXPROCS(4)  // 设置go程序只用一个逻辑核心
	go a()
	go b()
	wg.Wait()
}


var wg sync.WaitGroup