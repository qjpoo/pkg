package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

//为Person定义方法
func (p *Person) PrintInfo() {
	fmt.Printf("%s,%s,%d\n", p.name, p.sex, p.age)
}

type Student struct {
	Person
	id   int
	addr string
}

//Student定义方法，实际上就相当于方法重写
func (s *Student) PrintInfo() {
	fmt.Printf("Student:%s,%s,%d\n", s.name, s.sex, s.age)
}

func main() {
	p := Person{"接客", "male", 18}
	p.PrintInfo()
	//学生也去调，方法继承
	s := Student{Person{"接客", "male", 18}, 2, "bj"}
	s.PrintInfo()
	//显式调用
	s.Person.PrintInfo()
}