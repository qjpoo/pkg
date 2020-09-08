package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch1 chan int
var ch2 chan int

func main() {
	// select多路复用,可处理一个或多个channel的发送/接收操作

	ch1 = make(chan int,1024)
	ch2 = make(chan int,1024)
	go f1(ch1)
	go f2(ch2)

	for {
		select {
		case data := <-ch1:
			fmt.Println("ch1: ",data)
			//time.Sleep(time.Millisecond  * 200)
		case data := <-ch2:
			fmt.Println("ch2: ", data)
			//time.Sleep(time.Millisecond  * 500)
		default:
			fmt.Println("waiting ...")
			time.Sleep(time.Millisecond  * 100)
		}
	}





}

func f1(ch chan int)  {

	for i:=0;i <= rand.Int();i++ {
		ch1 <- i
		time.Sleep(time.Millisecond * 200)
	}
}

func f2(ch chan int)  {

	for i:=0;i< rand.Int();i++ {
		ch2 <- i
		time.Sleep(time.Millisecond * 500)
	}
}

