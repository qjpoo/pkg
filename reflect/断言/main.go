package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// 接口值
	var x interface{} // 定义了一个空接口,这个空接口里面没有方法, interface是关键字
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1

	x = a     // x接口的值由两部分组成, 一个是类型,一个是值  类型是int64 值是100
	x = b     // int32 10
	x = c     // int8 1
	x = 12.34 // float64, 12.34
	//x = false // <bool fase>
	fmt.Println(x)
	// 猜一猜, 类型断言

	value, ok := x.(bool)  // 猜x是bool类型
	// 如果猜对了,ok是true, value=对应类型的值
	// 如果猜错了,ok是false, value=对应类型(就是你这个括号里面写的类型)的零值
	if ok {
		fmt.Printf("%T, %v\n", value, value)  // 如果猜中了,value就是一个bool类型的值
	}else {
		fmt.Printf("%T, %v\n", value, value)  // 猜不中的话,value是bool类型, value值就是默认false
		//fmt.Println("x的值不是bool类型 ")

	}

	var w io.Writer // writer类型,它也是一个接口
	w = os.Stdout
	fmt.Printf("%T, %#v\n", w, w) // *os.File, &os.File{file:(*os.file)(0xc000072280)}

	w = new(bytes.Buffer)         // buffer是一个结构体类型
	fmt.Printf("%T, %#v\n", w, w) // *bytes.Buffer, &bytes.Buffer{buf:[]uint8(nil), off:0, lastRead:0}

	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))

}
