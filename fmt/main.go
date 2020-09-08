package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	// fmt
	var foo string = "This is a simple string"
	fmt.Printf("%v\n", foo)
	fmt.Printf("%T\n", foo)

	fmt.Println("---------------------")
	order := Person{
		"chiling",
		47,
	}
	// 这是我调试时的默认格式
	fmt.Printf("%+v\n\n", order)

	// 当我需要知道这个变量的有关结构时我会用这种方法
	fmt.Printf("%#v\n\n", order)

	// 我很少使用这些
	fmt.Printf("%v\n\n", order)
	fmt.Printf("%s\n\n", order)
	fmt.Printf("%T\n", order)

	fmt.Println("*******************************")
	ok, e := someFunction(0)
	fmt.Println("ok: ", ok, "e:　", e)

}

func someFunction(val int) (ok bool, err error) {
	fmt.Printf("Before val: %T, ok: %T, err: %T\n", val, ok, err)
	if val == 0 {
		fmt.Printf("After val: %T, ok: %T, err: %T\n", val, ok, err)
		return
	}
	if val < 0 {
		err = fmt.Errorf("value can't be negative %d", val)
		return
	}
	ok = true
	return
}
