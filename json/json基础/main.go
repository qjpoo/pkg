package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 序列化
	//var m = map[int]string{1:"a", 2: "b"}
	//fmt.Println(m[1])

	// 1. 字符串序列化之后还是字符串
	s := "hello, world"
	s1, err := json.Marshal(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s1))
	fmt.Println(s1)

	// 2. 数组的序列化
	a := [5]string{"a", "b", "c", "d"}
	a1, _ := json.Marshal(&a)
	fmt.Println(a1)
	fmt.Println("arry: ",string(a1))

	// map
	m := map[string]interface{}{"a":1.0, "b":"2b", "c":3}
	m1, _ := json.Marshal(&m)
	fmt.Println(string(m1))
	mm := make(map[string]interface{},10)
	json.Unmarshal(m1, &mm)
	fmt.Println(mm)
	for k, v := range mm {
		fmt.Println(k, "----->", v)
	}

}
