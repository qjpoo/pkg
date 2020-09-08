//package 声明开头表示代码所属包
package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

//定义传递学生对象的方法
func tmpStudent(tmp Student) {
	//修改id
	tmp.id = 250
	fmt.Println("tmp=", tmp)
}

//定义传递指针类型的对象的方法
func tmpStudent2(p *Student) {
	//修改id
	p.id = 249
	fmt.Println("tmp2=", p)
}

func main() {
	var s Student = Student{1, "接客", "female", 20, "sz"}
	tmpStudent(s)
	fmt.Println("main s =", s)
	//传递指针地址
	tmpStudent2(&s)
	fmt.Println("main s2=", s)
}

//tmp= {250 接客 female 20 sz}
//main s = {1 接客 female 20 sz}
//tmp2= &{249 接客 female 20 sz}
//main s2= {249 接客 female 20 sz}