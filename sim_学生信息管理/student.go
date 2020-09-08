package main

import "fmt"

type Student struct {
	name  string
	age   int
	id    int64
	class string
}

// 构造函数
func NewStudent(name string, age int, id int64, class string) *Student {
	return &Student{
		name:  name,
		age:   age,
		id:    id,
		class: class,
	}
}

// 定义一个学生信息管理的结构休
type StudentMgr struct {
	AllStudent []*Student
}

func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		AllStudent: make([]*Student, 0, 100),
	}

}

func (s *StudentMgr) AddStudent(stu *Student) {
	// 首先要判断StudentMgr里面有没有stu这个学生
	for _, v := range s.AllStudent {
		// 没有就添加
		if stu.name == v.name {
			fmt.Printf("添加出错,姓名为%s的学生已存在了 !\n", stu.name)
			return
		}
	}
	s.AllStudent = append(s.AllStudent, stu)
	//fmt.Println("s.AllStudent: ", s.AllStudent) //  &{[0xc0000663c0]}
}

// 编辑学生
func (s *StudentMgr) EditStudent(stu *Student) {
	// 首先要判断StudentMgr里面有没有stu这个学生
	for index, v := range s.AllStudent {
		if stu.name == v.name {
			s.AllStudent[index] = stu
			return
		}
	}
	fmt.Printf("编辑出错,姓名为%s的学生不存 !\n", stu.name)
}

// 展示学生
func (s *StudentMgr) ShowStudent() {
	//fmt.Println(s)
	for _, v := range s.AllStudent {
		fmt.Printf("学生姓名为: %s, 年龄为: %d, Id为: %d, 班级为: %s\n", v.name, v.age, v.id, v.class)
	}
}

// 删除学生
func (s *StudentMgr) DeleteStudent(stu *Student) {
	for index, v := range s.AllStudent {
		if stu.name == v.name {
			// 从切片按照索引删除指定的元素
			// ...把后面的切片,拆开了一个一个的添加进去
			s.AllStudent = append(s.AllStudent[:index], s.AllStudent[index+1:]...)
			fmt.Println("-----------------------------")
			s.ShowStudent()
			return
		}
	}
	// for循环完了,还没有找到相等的,说明没有这个用户
	fmt.Printf("删除出错,姓名为%s的学生不存 !\n", stu.name)
}
