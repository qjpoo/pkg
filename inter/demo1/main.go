package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
type xxoo interface {
	call()
}

func call(x xxoo)  {
	fmt.Println(x.call())

}
 */

type caller interface {
	call()
}

func ko(x caller)  {
	fmt.Printf("%T, %v\n", x, x)
	x.call()
}
type animal struct {
	name string
	age int
}

type duck struct {
	name string
	age int
}

type chicken struct {
	name string
	age int
}

type monkey struct {
}

func (a animal) call()  {
	fmt.Println(a.name, "animal call ...")
}

func (d duck) call()  {
	fmt.Println("duck call ...")
}

func (c chicken) call()  {
	fmt.Println("chicken call ...")
}

func (c monkey) call()  {
	fmt.Println("monkey call ...")
}

func main() {
	// interface value

	var w io.Writer
	fmt.Printf("%T, %v\n", w, w)  // <nil>, <nil>

	var xxoo interface{}
	fmt.Printf("%T, %v\n", xxoo, xxoo)  // <nil>, <nil>

	w1 := new(bytes.Buffer)
	fmt.Printf("%T, hello_server %v hello_server\n", w1, w1)  // *bytes.Buffer, hello_server  hello_server

	//n, err := w1.Write([]byte("hello_server")) // 正常
	//n, err := w.Write([]byte("xxoo"))  // panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(n, err)

	var a1 animal = animal{"animal one", 20}
	ko(a1)

	var d1 duck
	ko(d1)

	var c1 = chicken{name: "chicken1", age: 1}
	fmt.Println(c1)
	fmt.Printf("%T, %v\n", c1, c1)

	var c2 caller
	fmt.Printf("%T, %v\n", c2, c2)

	var c3 caller = c1
	fmt.Printf("%T, %v\n", c3, c3)

	var m monkey
	fmt.Printf("%T, %v\n", m, m)

	// 取一个变量的指针
	go1 := 100
	p := &go1
	fmt.Println(go1,p,*p,&p)
}
