package main

import "fmt"

type name int8

type first struct {
	a int
	b bool
	name
}
func main() {

	var a first = first{1, false, 2}
	var b *first = &a
	fmt.Printf("%T, %v, %v, %v\n", b, b, *b, &b)  //  *main.first, &{1 false 2}, {1 false 2}, 0xc000006028
	fmt.Printf("%T, %v, %v, %v\n", a, a.a, a.b, &a)  // main.first, 1, false, &{1 false 2}
	fmt.Println(a,b, a.b, a.a, a.name, &a, b.a, &b, (*b).a)  /*
	a {1 false 2}
	b &{1 false 2} 指针是地址
	a.b false
	a.a 1
	a.name 2
	&a    &{1 false 2}   a是结构体变量  &a就是 &{1 false 2}
	b.a    1
	b   &{1 false 2}
	&b 0xc000006028   获取指针地址在指针变量前加&的方式
	(*b).a 1

	*/
	var x int = 10
	var p *int
	p = &x
	fmt.Printf("%p, %p", p, p)

}
