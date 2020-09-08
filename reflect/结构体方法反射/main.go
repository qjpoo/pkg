package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" xx:"mingzi"`
	Score int    `json:"score" xx:"fengshu"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func (s student) Test(i int) int {
	fmt.Println("i = ", i)
	return i
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod()) // 2
	for i := 0; i < v.NumMethod(); i++ {
		// v.Method(i) 结构体的方法
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name) // method name:Sleep
		fmt.Printf("method:%s\n", methodType)            // method:func() string  方法没有传入参数

		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		//var args = []reflect.Value{}
		//v.Method(i).Call(args)

		arg := []reflect.Value{reflect.ValueOf(100)}
		v.MethodByName("Test").Call(arg)
	}
}

func main() {
	// 结构体方法的反射

	var stu = student{"chiling", 100}
	printMethod(stu)
}
