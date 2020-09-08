package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 反射
	// typeof 类型    valueof 值
	/*
	a := 100
	fmt.Println(reflect_typeof.TypeOf(a))  // int
	reflectType(a)  // type: int
	reflectType(true)  // bool
	reflectType([...]int{1,2,3}) // [3]int
	reflectType("xxoo") // string
	reflectType(map[string]int{})  //  map[string]int

	c1 := Cat{"chiling"}
	fmt.Println(reflect_typeof.TypeOf(c1))
	 */

	c1 := Cat{"chiling"}
	reflectType(c1)  // name: cat  kind: struct    kind是大范围的, name是指具体的

	i := 100
	reflectType(&i)  //name:     kind:  ptr  对于指针类型,它的name是空

	s := []int{1,2,3,4}
	reflectType(s)  //  name:  kind: slice

	m := map[string]int{"a": 1, "b":2}
	reflectType(m)   // name:  kind: map



}

// 反射
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	//fmt.Printf("type: %v\n", v)
	//fmt.Printf("%T\n", v)  // printf知道传进来的变量的类型,就是使用了反射, 代码实例也是使用的的反射
	fmt.Printf("name: %v kind: %v\n", t.Name(), t.Kind())

}
type Cat struct {
	name string
}

type Person struct {
	name string
	age int
}
