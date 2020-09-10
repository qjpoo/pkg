package main

import (
	"fmt"
	"os"
	"strconv"
)

func f1(args ...interface{}) {
	f2(args...)
	f2(args[1:]...)
}

func f2(args ...interface{}) {
	for i, v := range args {
		fmt.Fprintf(os.Stdout, "i = %d %v\n", i, v)
	}
	fmt.Fprintf(os.Stdout, "--------------\n")
}

func main() {
	f1(1, "hello", 3.14, main)

	// 匿名函数 1
	f := func(i, j int) (result int) { // f 为函数地址
		result = i + j
		return
	}

	fmt.Fprintf(os.Stdout, "f = %v  f(1,3) = %v\n", f, f(1, 3))

	// 匿名函数 2
	x, y := func(i, j int) (m, n int) { // x y 为函数返回值
		return j, i
	}(1, 9) // 直接创建匿名函数并执行

	fmt.Fprintf(os.Stdout, "x = %d   y = %d\n", x, y)

	fmt.Fprint(os.Stdout, "-------------------->\n")
	stringToInt("123", myLog)
	//stringToInt("s", myLog)

	n1 := Test{
		param1: struct{ param2 string }{param2: "TEST"},
	}

	fmt.Println(n1)

	n2 := Test1{
		param1: struct {
			param2 string `json:"param2"`
		}{param2: "test1"}}
	fmt.Println(n2)

	n3 := new(Test1)
	n3.param1.param2 = "xxoo"
	fmt.Println(n3)

	s1 := make([]struct{}, 10)
	s2 := make([]struct{}, 20)
	fmt.Println(&s1 == &s2)  // false
	fmt.Println(&s1[0] == &s2[0]) // true
}

type saveLog func(msg string)

//这个函数的第二个参数是一个函数
//这个函数将一个字符串转换为Int类型，如果失败了，则返回0，并输出错误。
func stringToInt(s string, log saveLog) int64 {
	if value, err := strconv.ParseInt(s, 0, 0); err != nil {
		log(err.Error())
		return 0
	} else {
		return value
	}
}

//记录日志的函数实现
func myLog(msg string) {
	fmt.Println("Find error:", msg)
}

type Test struct {
	param1 struct {
		param2 string
	}
}

type Test1 struct {
	param1 struct {
		param2 string `json:"param2"`
	}
}
