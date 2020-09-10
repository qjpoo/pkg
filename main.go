package main

import (
	"fmt"
)

func main() {
	x := [3]int{1, 2, 3}
	x1 := &x
	fmt.Println(x1)
	func(arr *[3]int) {
		fmt.Println(arr)
		(*arr)[0] = 7
		fmt.Println(arr) // &[7 2 3]
	}(&x)
	fmt.Println(x) // [7 2 3]

	var data interface{} = "great"

	if res, ok := data.(int); ok {
		fmt.Println("[is an int], data: ", res)
	} else {
		fmt.Println("[not an int], data: ", data) // [not an int], data:  great
	}

	/*
	m3 := map[string]data1{
		"x": {"Tom"},
	}
	m3["x"].name = "Jerry"
	 */

	y1 := []data1{{name: "xxoo"},{name: "xxaa"}}
	y1[0].name="xx11"
	fmt.Println(y1)


}

type data1 struct {
	name string
}
