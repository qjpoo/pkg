package main

import (
	"encoding/json"
	"fmt"
)

// student
type Student struct {
	ID int
	Gender string
	Name string
}

// class
type Class struct {
	Title string
	Student []*Student
}



func main() {
	c := &Class{
		Title: "101",
		Student: make([]*Student, 0, 200),
	}
	//fmt.Printf("%T", c)

	for i:=0;i<10;i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%02d", i),
			Gender:  "男",
			ID: i,
		}
		c.Student = append(c.Student,stu)
	}

	// json序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json failure ...")
		return
	}

	fmt.Printf("json: %s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	fmt.Printf("%T\n", err)
	if err != nil {
		fmt.Println("json unmarshal failure ...")
		return
	}
	fmt.Println(c1)


}
