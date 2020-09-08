package main

import (
	"fmt"
	"time"
)

type xxoo struct {
	name string
}

func main() {
	queue := make(chan int, 1)

	go Producer(queue)

	go Consumer(queue)

	time.Sleep(1* time.Second) //让Producer与Consumer完成
	//time.Now()
	//fmt.Println()

	var p * int
	var i int
	i = 10
	p = &i
	fmt.Println(p)  // 0xc0000160a8

	//var xo xxoo = xxoo{"hello_server"}
	xo := new(xxoo)
	p1 := &xo
	fmt.Println(xo, &xo, p1, (*p1).name, "xxoo")

}



func Producer(queue chan<- int) {

	for i := 0; i < 10; i++ {

		queue <- i

	}

}

func Consumer(queue <-chan int) {

	for i := 0; i < 10; i++ {

		v := <-queue

		fmt.Println("receive:", v)

	}

}
