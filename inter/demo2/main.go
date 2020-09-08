package main

import "fmt"

type animal interface {
	speak()
}

type cat struct {
	name string
}


//func (c cat) speak() {
//	fmt.Printf("喵喵 ... value")
//}

func (c *cat) speak() {
	fmt.Printf("喵喵 ... address ")
}

func main() {
	var x animal
	//c1 := cat{name: "小猫"}
	//x = c1
	//x.speak()

	c2 := cat{"小猫"}
	x = &c2
	x.speak()

	c4 := cat{"小猫"}
	x = c4
	x.speak()

	c3 := cat{"小猫"}
	c3.speak()



}
