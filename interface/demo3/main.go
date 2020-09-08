package main

import "fmt"

type A interface {
	Print()
}

type B struct {

}

func (b *B)Print()  {
	fmt.Println("bbbb ...")

}


func main() {
	var a A
	a = &B{}
	a.Print()

	var aa []A
	aa = make([]*B, 10)



}
