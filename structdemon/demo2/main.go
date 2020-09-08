package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。

注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。
*/
type Student struct {
	//ID     int `json:"id"`
	ID     int `json:"id"`
	Gender string
	name   string
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "小王",
	}

	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failure ...")
		return
	}
	fmt.Printf("json: %s", data)

}
