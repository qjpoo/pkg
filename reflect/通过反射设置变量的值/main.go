package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 通过反射修改值

	var a int64 = 100
	// 将a传入一个函数,在函数中修改a的值 ,该函数接受任意的值
	modifyValue(&a)  // 函数传参永远是值COPy
	fmt.Println(a)

	x := 1
	f(&x)
	fmt.Println(x)

	var y interface{}
	fmt.Printf("%T, %v\n", y,y)  // <nil>, <nil>

	var xx Person
	fmt.Printf("xx: %T, xx: %v\n", xx, xx)  // xx: <nil>, xx: <nil>
	xx = &Student{}
	fmt.Printf("%T, %v\n", xx,xx)  // *main.Student, &{}

	//n := Student{"chiling"}
	var n  Person
	Check(n)


}

func Check(test interface{})  {
	//_, ok := test.(Student)
	//v, ok := test.(Student)
	v, ok := test.(Person)
	fmt.Printf("v type: %T, v value: %v\n", v, v)
	fmt.Println(ok)

}


func modifyValue(a interface{}) {
	v := reflect.ValueOf(a)  // reflect.value
	kind := v.Kind()
	fmt.Println(kind)  // ptr
	fmt.Println("v.Elem().Kind(): ", v.Elem().Kind())  // int64
	// 反射中使用elem()方法获取指针对应的值
	if kind == reflect.Ptr {
		// 传入的是一个指针
		// 取到指针对应的值再去修改
		v.Elem().SetInt(200)
	}
}

func f(x *int)  {
	*x=123

}

type Person interface {
	Say()
}

type Student struct {
	Name string
}

func (s *Student)Say()  {
	fmt.Println(s.Name)
}
