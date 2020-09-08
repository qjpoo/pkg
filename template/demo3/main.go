package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Say() Person  {
	fmt.Println(p.name, "Say ...")
	return p
}

func (p Person) Move() Person  {
	fmt.Println(p.name, "Move ...")
	return p
}
func main() {
	//链式操作
	// 原理: 每一次执行完成之后返回的是操作的对象本身

	var p = Person{
		name: "chiling",
	}
	p.Say().Move()

}
