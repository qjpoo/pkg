package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

func main() {
	p := Person{"luhan", "男"}
	//1.生成JSON文本
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//2.生成格式化json，没啥用
	b, err = json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
