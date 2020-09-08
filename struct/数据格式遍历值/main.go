package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	txt := `{
		"a":1,
		"b":2,
		"c":[
			{"name":"1","group":"2"},
			{"name":"3","group":"4"}
		]
	}`
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(txt), &m); err != nil { // 反序列化 json ==> map[string]interface{}
		panic(err)
	}
	fmt.Println([]byte(txt))
	fmt.Println(m)
	v := reflect.ValueOf(m["c"])
	fmt.Println("v: ", v)
	count := v.Len()
	fmt.Println(count)
	for i := 0; i < count; i++ {
		fmt.Println(v.Index(i))
	}
}
