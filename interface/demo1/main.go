package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Swap(i, j int) {
	// 实际走的交换逻辑
	p[i], p[j] = p[j], p[i]
}

// 一般情况下，上面这俩不需要变化，但是我们实际比较的可能是不同的维度

type SortByName struct{ Persons }

// 继承了Persons上面的方法

func (p SortByName) Less(i, j int) bool {
	// 根据元素的姓名长度排序
	return len(p.Persons[i].Name) > len(p.Persons[j].Name)
}

type SortByAge struct{ Persons }

func (p SortByAge) Less(i, j int) bool {
	return p.Persons[i].Age > p.Persons[j].Age
}

func main() {
	persons := Persons{
		{
			Name: "tes",
			Age:  20,
		},
		{
			Name: "test",
			Age:  22,
		},
		{
			Name: "test3asdf",
			Age:  21,
		},
	}

	fmt.Println("排序前")
	for _, person := range persons {
		fmt.Println("name", person.Name, "age", person.Age)
	}
	sort.Sort(SortByName{persons})
	fmt.Println("排序后")
	for _, person := range persons {
		fmt.Println("name", person.Name, "age", person.Age)
	}

}
