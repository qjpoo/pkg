package main

import "fmt"

var ch1 chan int
var ch2 chan int

func c1(ch1, ch2 chan int)  {
	ch1 <- 10
	ret := <- ch1
	ch2 <- ret
}

func main() {
	// ch1, ch2之关传递数据
	ch1 = make(chan int, 10)
	ch2 = make(chan int, 10)
	c1(ch1, ch2)
	ret := <- ch2
	fmt.Println(ret)
}
