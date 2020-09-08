package main

import "fmt"

type Human interface {
	Len()
}

type Student interface {
	Human
}

type Test struct {
	//age int
	//yes bool
}

func (h *Test) Len()  {
	fmt.Println("成功")
}


func main() {
	var s Student
	fmt.Printf("%T, %v\n", s, s)
	// s.Len() // panic: runtime error: invalid memory address or nil pointer dereference
	s = new(Test)
	fmt.Printf("%T, %v, %v\n", s, s, &s)   // *main.Test, &{}, 0xc00003a1f0
	s.Len()

	s = &Test{}
	fmt.Printf("%T, %v, %v\n", s, s, &s)

	var h Human
	h = new(Test)
	h.Len()

	var h1 Human
	h1 = &Test{}
	h1.Len()



}
