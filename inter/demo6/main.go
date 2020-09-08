package main

import (
	"bufio"
	"fmt"
	"io"
)

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

func abc() string {
	return "hello_server"
}
type Cat struct { // 如果想实现接口C，那么就需要把C中的接口A,B两个接口的方法也要实现了，还要实现C中的test3方法
	abc func() string
}

func (c Cat) test1()  {
	fmt.Println("test1()")
}

func (c Cat) test2()  {
	fmt.Println("test2()")
}

func (c Cat) test3()  {
	fmt.Println("test3()")
}


func main() {
	// 接口的继承

	// 同样是一个Cat对象，那它到底能访问那一个方法，就看你给它定义的是哪一个类型，比如你定义民A, B, C
	var cat = Cat{}
	fmt.Printf("%T, %v\n", cat, cat)
	cat.test1()
	cat.test2()
	cat.test3()
	//cat.abc()

	//xxoo := new(Cat)
	//fmt.Printf("%T, %v\n", xxoo, xxoo)
//	fmt.Println(xxoo.abc())
	//n := xxoo.abc()
	//fmt.Println(n)

	//xxoo1 := Cat{abc:abc}
	//fmt.Printf("%T, %v\n", xxoo1, xxoo1)
	//fmt.Println(xxoo1.abc())


	fmt.Println("------------------------")
	var a1 A = Cat{}
	//fmt.Printf("%T, %v", a1, a1)
	a1.test1() // 定义了a1是A接口，只能调用A接口的test1方法


	var a2 B = Cat{}
	a2.test2()

	var a3 C = Cat{}
	a3.test2()
	a3.test1()
	a3.test3()


	//a1 = a2

	var n1 A
	fmt.Printf("%T, %v", n1, n1)

	var n2 B
	fmt.Printf("%T, %v", n2, n2)

	var cat1 Cat
	fmt.Printf("%T, %v\n", cat1, cat1)

	fmt.Println("==------------")
	var w1 io.Writer // interface
	fmt.Printf("%T, %v\n", w1, w1)
	fmt.Println("==------------")
	x1 := new(bufio.Writer) // struct
	fmt.Println(x1)
	x1 = nil
	fmt.Printf("%T, %v", x1,x1)

	//bytes.Buffer{}



}
