package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("%T, %v, %v\n", a, a, &a)  // int, 10, 0xc0000120b8

	// *T是指针变量的类型，它指向T类型的值
	var p *int  // 声明指针变量
	p = &a  // 指针变量的存储地址
	fmt.Printf("%T, %v, %v, %v\n", p, p, &p, *p)  // *int, 0xc0000120b8, 0xc000006030, 10

}
