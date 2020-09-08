package main

import (
	"fmt"
	"reflect"
)

//定义一个类型
type tsh struct {
	//定义成员,类型是func() string
	test func() string
}

//定义一个函数,获取tsh类型
func New(fn func() string) *tsh {
	fmt.Println(fn, reflect.TypeOf(fn))  // 0x49fb50 func() string
	return &tsh{
		test: fn,
	}
}
func cre() string {
	return fmt.Sprintf("%s,来了", "tsh")
}

func main() {
	//new完得到tsh类型,调用该结构体的test成员,该成员是个函数
	res := New(cre).test()
	fmt.Println(New(cre))  // &{0x49fb50}
	fmt.Println(res)
}
