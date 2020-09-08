package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//1.准备一段json
	b := []byte(`{"Name":"luhan","Hobby":"男"}`)
	//2.声明空接口
	var i interface{}
	//3.解析
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println("json err:", err)
	}
	//自动转map
	fmt.Printf("%T, %#v\n",i, i)
	//4.使用interface的json，可以判断类型
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "是string类型", vv)
		case int:
			fmt.Println(k, "是int类型", vv)
		default:
			fmt.Println(k, "是其他类型")
		}
	}
}
