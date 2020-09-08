package main

import "fmt"

type Person struct {
	name string
	age int
}

func tell(person Person)  {
	fmt.Printf("hi, my name's %v, i'm %v years old ...\n", person.name, person.age)
}


func (person Person) tell() {
	fmt.Printf("Hi, my name's %v, I'm %v years old\n", person.name, person.age)
}

func main() {

	fmt.Println("------------------函数-------------------")
	// 调用函数，提供一个struct实例
	person := Person{"Tom", 28}
	tell(person) // Hi, my name's Tom, I'm 28 years old


	// 调用函数，提供一个struct实例的地址
	//person := Person{"Tom", 28}
	//tell(&person) // 编译错误：Cannot use '&person' (type *Person) as type Person

	fmt.Println("--------------方法体---------------")
	person1 := Person{"Tom", 28}
	person1.tell() // Hi, my name's Tom, I'm 28 years old

	// 得到一个结构体指针
	var person2 *Person = &Person{"Tom", 28}
	person2.tell() // Hi, my name's Tom, I'm 28 years old

}
