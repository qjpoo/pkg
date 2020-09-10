package main

import (
	"fmt"
)

//解析(map[string]interface{})数据格式并打印出数据

func print_json(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float", int64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_json(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}

func main() {

	mapInterface := make(map[interface{}]interface{})
	mapString := make(map[string]string)

	mapInterface["k1"] = 1
	mapInterface[3] = "hello"
	mapInterface["world"] = 1.05

	for key, value := range mapInterface {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapString[strKey] = strValue
	}

	fmt.Printf("%#v", mapString)
}
