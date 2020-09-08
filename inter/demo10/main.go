package main

import "fmt"

func main() {
	// 空接口

	var a1 A = Cat{"黑猫"}
	var a2 A = Person{"小明", 20}
	var a3 int = 100
	var a4 bool = true
	var a5 string = "hello_server"
	fmt.Println(a1, a2, a3, a4, a5)

	var a6 A
	fmt.Println(a6) // nil

	// 调用空接口
	nilIf("string")
	nilIf(10)
	nilIf(nil)

	//
	nilIf2(3.15)

	// map
	map1 := make(map[string]interface{})
	map1["1"] = "quinn"
	map1["3"] = 1
	map1["2"] = 2.1
	map1["4"] = Person{name: "chiling", age: 100}
	map1["5"] = [...]int{1, 2, 3}
	for k, v := range map1 {
		fmt.Println(k,"--",v)
	}

	// slice
	fmt.Println("--------------")
	s := make([]interface{},4, 10)
	s[0] = 1
	s[1] = 1.0
	s[2] = "string"
	s[3] = 'A'
	for k, v := range s {
		fmt.Println(k, "----", v)
	}

}

// 空接口
type A interface {
}
type Cat struct {
	color string
}

type Person struct {
	name string
	age  int
}

// 函数的参数是一个空接口
func nilIf(a A) {
	fmt.Println(a)
}

func nilIf2(a interface{}) {
	fmt.Println(a)
}
