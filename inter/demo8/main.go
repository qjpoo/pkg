package main

import (
	"fmt"
	"math"
)

type Shape interface {
	peri() float64
	area() float64
}

// 三角形
type triangle struct {
	a, b, c float64
}

// 圆
type circle struct {
	r float64
}

func (t triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t triangle) area() float64 {
	p := t.peri() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (c circle) peri() float64 {
	return math.Pi * 2 * c.r
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

// 三角形的构造函数
func NewTriangle(a, b, c float64) *triangle {
	return &triangle{
		a: a,
		b: b,
		c: c,
	}
}

// 得到形状数据的函数
func GetData(s Shape) { // 此时这里不能得到s的具体类型，比如说是三角形还是圆形
	fmt.Printf("peri: %f\tared: %f\n", s.peri(), s.area())
}

// 得到具体的形状
func GetType(s Shape) {
	/*
		ins, ok := s.(triangle)  // {0 0 0} true , s接口的形状如果是三角形的话，ok就为真 ins就是返回状态的实例
		fmt.Println(ins, ok)
		if ok  { // if ins, ok := s.(triangle);ok {
			fmt.Println("triangle find ...", ins.a, ins.b, ins.c)
		}

		if ins1, ok1 := s.(circle);ok1 {
			fmt.Println("circle find ...", ins1.r, ok1)
		}

		if ins2, ok2 := s.(Shape); ok2 {
			fmt.Println("shape find ...", ins2.area(), ins2.peri())
		}
	*/

	/*
		if ins, ok := s.(*triangle); ok {
			fmt.Printf("%T, %v, %v\n", ins, ins, &ins) // *main.triangle, &{3 4 2}, 0xc000006038
			fmt.Printf("%T, %v, %v\n", s, s, &s)  // *main.triangle, &{3 4 2}, 0xc00003a1f0
			fmt.Println("triangle find ...", ins.a, ins.b, ins.c)

		}
	*/

	switch ins := s.(type) {
	case triangle:
		fmt.Println("triangle  ", ins.a, ins.b, ins.c)
	case *triangle:
		fmt.Println("*triangle  ", ins.a, ins.b, ins.c)
	case circle:
		fmt.Println("circle  ", ins.r)
	default:
		fmt.Println("Not Found ...")

	}
}

func main() {
	// 接口断言

	t := NewTriangle(3, 4, 5)
	fmt.Println(t.a, t.b, t.c, t.area(), t.peri())

	// 申明一个圆对象
	var c circle = circle{5}
	fmt.Println(c.peri())
	fmt.Println(c.area())

	// 申明一个接口类型,接口是不能访问实现类中的字段的，比如s接口不能访问此时的三角形中的a , b , c三条边长
	var s Shape
	s = NewTriangle(3, 4, 5)
	s.peri()
	s.area()

	// 把圆形赋值给s接口
	s = c

	//
	//var t2 triangle = triangle{3,4,5}
	//GetData(t2) // 此时这里的参数，你可以传三角形， 圆形，或者 shape都是可以的

	//var t3 triangle
	//GetType(t3)

	//var c1 circle
	//GetType(c1)

	// 如果传的是指针地址
	var t4 *triangle = &triangle{3, 4, 2}
	fmt.Printf("%T, %v, %v\n", t4, t4, &t4) // *main.triangle, &{3 4 2}, 0xc00000c3e0
	GetType(t4)

}
