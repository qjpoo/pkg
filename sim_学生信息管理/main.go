package main

import (
	"fmt"
	"os"
)


func main() {
	// 学生信息管理
	/*
		学生信息管理
		命令行菜单
		修改学生
		删除学生
		展示学生
	*/


	stuMgr := NewStudentMgr()
	for {
		showMenu()
		var i int
		fmt.Print("请选择序号: ")
		fmt.Scanln(&i)
		switch i {
		case 1:
			fmt.Println("您选择的是: 1. 添加学生")
			stu := userInput()
			fmt.Println("stu: ",stu) // &{1 1 1 1}
			//fmt.Println("1 stuMgr: ", stuMgr)
			//stuMgr.AddStudent(stu)
			stuMgr.AddStudent(stu)
			//fmt.Println("2 stuMgr: ", stuMgr)
		case 2:
			fmt.Println("您选择的是: 2. 修改学生")
			stu := userInput()
			stuMgr.EditStudent(stu)
		case 3:
			fmt.Println("您选择的是: 3. 删除学生")
			stu := userInput()
			stuMgr.DeleteStudent(stu)
		case 4:
			fmt.Println("您选择的是: 4. 展示学生")
			//fmt.Println(stuMgr)
			stuMgr.ShowStudent()
		case 5:
			fmt.Println("您选择的是: 5. 退出系统")
			os.Exit(0)
		default:
			fmt.Println("选择有误 ...")
			os.Exit(0)
		}

	}

}

// 信息显示
func showMenu() {
	fmt.Println("学生信息管理系统 !")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 修改学生")
	fmt.Println("3. 删除学生")
	fmt.Println("4. 展示学生")
	fmt.Println("5. 退出系统")
}

func userInput() *Student {
	var (
		name  string
		age   int
		id    int64
		class string
	)
	fmt.Println("请按提示输入信息: ")
	fmt.Printf("请输入姓名: ")
	fmt.Scanln(&name)
	fmt.Printf("请输入年龄: ")
	fmt.Scanln(&age)
	fmt.Printf("请输入序号: ")
	fmt.Scanln(&id)
	fmt.Printf("请输入班级: ")
	fmt.Scanln(&class)

	stu := NewStudent(name, age , id, class)
	return stu

}
