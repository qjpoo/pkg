package main

import "fmt"

func main() {
	var a int
	var b int
	a, b = 1, 2
	fmt.Println("before  a, b : ", a, b)
	change(a, b)
	fmt.Println("after  a, b : ", a, b)
	pointchange(&a , &b)
	fmt.Println("after  a, b : ", a, b)
}

func pointchange(a, b *int)  {
	*a, *b = *b, * a
	fmt.Println("pointchange: ",*a, *b)

}
func change(a, b int)  {
	b, a = a, b
	fmt.Println(a, b)
}
