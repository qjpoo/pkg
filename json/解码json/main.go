package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func main() {
	//1.准备一段json
	//b := []byte(`{"NAME":"luhan","Hobby":"男"}`)
	fmt.Printf("%#v\n", string(b))
	b := []byte(`{"name":"luhan","hobby":"男"}`)
	//2.声明结构体
	var p Person
	//3.解析
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(p)
}
