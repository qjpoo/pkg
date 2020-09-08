package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float64
}

type Employee struct {
	Human
	company string
	money float64
}

func (h Human) Sayhi()  {
	fmt.Println("hi, ", h.name)
}

func (h Human) Sing(song string)  {
	fmt.Println("La la ..., ", song)
}

// employee重写了Human的Sayhi方法
func (e Employee) Sayhi()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// men interface被Human， student, employee结构体三个实现了
type Men interface {
	Sayhi()
	Sing(song string)
}



func main()  {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	//paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	//sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	//Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	// 定义一个Men类型的i
	var i Men
	fmt.Printf("%T, %v\n", i, i)
	i = mike
	fmt.Printf("%T, %v\n", i, i)  // main.Student, {{Mike 25 222-222-XXX} MIT 0}
	test := Student{}
	fmt.Printf("%T, %v\n", test, test)  // main.Student, {{ 0 }  0}

}
