package main

import "fmt"

type name int8
type first struct {
	a int
	b bool
	name
}

func main() {
	a := new(first)
	fmt.Printf("%T, %v\n", a, a)
	a.a = 1
	a.name = 10
	fmt.Println(a.b, a.a, a.name)  // false 1 10

}