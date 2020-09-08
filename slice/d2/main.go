package main

import "fmt"

func printAll(vals []interface{}) { //1
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	//names := []string{"stanley", "david", "oscar"}

	//这个会报错
	//printAll(names)
	var dataSlice = []string{"a", "b", "c"}
	var interfaceSlice = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	printAll(interfaceSlice)
}
