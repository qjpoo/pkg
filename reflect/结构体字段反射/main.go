package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" xx:"mingzi"`
	Score int    `json:"score" xx:"fengshu"`
}

func main() {
	// 结构体反射示例

	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	//fmt.Printf("%T, %v\n", t, t)  //  *reflect.rtype, main.student
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//fmt.Printf("%T, %v\n", field, field)  // reflect.StructField, {Name  string json:"name" 0 [0] false}
	//	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("xx"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
