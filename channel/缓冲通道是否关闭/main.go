package main

import "fmt"

func send(ch chan int)  {
	for i:=0;i<10;i++ {
		ch <- i
	}
	close(ch)

}
func main() {
	// 判断通道是否关闭

	var ch1 chan int
	ch1 = make(chan int, 100)
	go send(ch1)
	/*
	// 通过for循环拿cha里面的值
	for {
		ret, ok := <-ch1
		if !ok { // 当通道关闭了,ok为false
			//fmt.Println(ok)  // false
			break
		}
		fmt.Println(ret)
	}

	 */

	// 利用for range循环去通道ch1中接收值
	for ret := range ch1 {
		fmt.Println(ret)
	}

}
