package main

import "fmt"

func output1()  {
	var (
		name string
		age   int
	)
	fmt.Printf("请输入name, age: ")
	fmt.Scanln(&name, &age)
	fmt.Println()
	fmt.Printf("扫描结果 ： name: %s, age: %d\n", name, age)
}
func main() {
	output1()
}
