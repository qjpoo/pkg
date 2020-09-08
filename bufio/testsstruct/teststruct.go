package main

import "fmt"

type student struct {
	name string
	age  int
}

// 构造函数
func newStudent(name string, age int)  *student {
	return &student{
		name: name,
		age: age,
	}
}


func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八的儿子", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		//fmt.Printf("%v, %v, %v\n", &stu, stu.name, stu)
		fmt.Printf("%v\n", &stu)
	}
	fmt.Println("---------------------")
	fmt.Println(m["小王子"])
	for k, v := range m {
		//fmt.Println(k, "=>", v.age)
		fmt.Println(k, "=>", v.name)
	}


	//
	var s1 student = student{"quinn", 20}
	fmt.Println(s1)

	// 构造函数
	// 构造函数 ，是一种特殊的方法。主要用来在创建对象时初始化对象， 即为对象成员变量赋初始值，总与new运算符一起使用在创建对象的语句中
	var s2 *student
	s2 = newStudent("quinn", 35)
	fmt.Println(s2)


}
