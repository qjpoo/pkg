package main

import "fmt"

func main() {

	b := 255
	a := &b
	fmt.Printf("%T, %T, %v, %v\n", b,a,b,a)  // int, *int, 255, 0xc0000120b8
	fmt.Printf("%v, %v\n", &b, *a)
	*a++
	fmt.Println(*a)
}
