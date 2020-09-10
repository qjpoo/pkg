package  main

import "fmt"

type Tester interface {
	Do()
}

type FuncDo func()
func (self FuncDo) Do() {
	self()
	fmt.Println("xxoo")
}
func main() {

	var t Tester = FuncDo(func() {
		fmt.Println("hello, world ...")
	})
	t.Do()
}
