package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {


	wg.Add(10)
	for i:=0;i<10;i++ {
		go hello(i)
	}
	fmt.Println("main func ...")
	wg.Wait()
}

func hello(i int)  {
	fmt.Println("hello_server func .... i: ", i)
	wg.Done()
}
