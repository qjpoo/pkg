package main

import (
	"bytes"
	"fmt"
	"io"
)

func test(w io.Writer)  {
	fmt.Println("w: ", w, w !=nil) // 接口有两个值 ，一个是类型，一个是接口的值，两个有一个不为空，就不为空
	if w != nil {
		w.Write([]byte("ok"))
	}
	
}

func main() {
	// 在golang中，接口值是由两部分组成的，一部分是接口的类型，另一部分是该类型对应的值，我们称其为动态类型和动态值。
	// 在第一行定义变量w的时候，声明了其类型为io.Writer，这里是真正意义上的空接口.
	//为什么是空接口，就是它的类型和值都为nil，在这里可以用==或者!=来和nil做判断。
	var w io.Writer // interface
	fmt.Printf("%T, %v, %p\n", w, w, &w)  // <nil>, <nil>, 0xc00003a220

	// 在第二行为变量w赋值的时候，此时w的动态类型为*bytes.Buffer，然后动态值是一个指向新分配的缓冲区的指针。
	w = new(bytes.Buffer) // struct  bytes.Buffer实现了这个接口
	fmt.Printf("%T, %v, %p\n", w, w, &w)  //  *bytes.Buffer, , 0xc000006030

	w = nil
	fmt.Printf("%T, %v, %p\n", w, w, &w)  //  <nil>, <nil>, 0xc00003a220

	//fmt.Println("----------*bytes.Buffer---------------")
	//var buf *bytes.Buffer // panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Printf("%T, %v\n", buf, buf)  // *bytes.Buffer, <nil>
	//test(buf)


	//buf1 := new(bytes.Buffer)
	//fmt.Printf("%T, %v\n", buf1, buf1)  // *bytes.Buffer,
	//test(buf1)
}
