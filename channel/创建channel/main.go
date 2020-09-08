package main

import (
	"fmt"
)

func main() {
	// 创建channel

	// 定义一个channel,它传递的是int类型的数据
	var ch1 chan int
	var ch2 chan string
	// channel是引用类型
	fmt.Println(ch1, ch2) // <nil> <nil>

	// make 函数初始化(分配内存)   slice map channel
	var a1 []int
	fmt.Println(a1) // []  没有初始化
	a1 = make([]int, 8)
	fmt.Println(a1) // 初始化  [0 0 0 0 0 0 0 0]

	var a2 = []int{}
	fmt.Println(a2) // []


	ch3 := make(chan int, 1)

	// 通道的操作, 发送 接收 关闭
	// 发送和接收  <-

	// 把10发送到ch3里面  发送
	ch3 <- 10
	ret := <- ch3
	fmt.Println(ret)
	ch3 <- 20

	// 从ch3接收值,保存到变量中
	 //<- ch3 丢弃了
	ret = <- ch3
	fmt.Println(ret)

	// 关闭
	close(ch3)

	// 关闭的通道再接收, 能取到对应类型的零值
	ret2 := <- ch3
	fmt.Println(ret2)  // 0

}
