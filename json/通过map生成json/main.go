package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建map
	mmp := make(map[string]interface{})
	mmp["name"] = "luhan"
	mmp["age"] = 56
	mmp["niubility"] = true

	//map转为json
	mjson, err := json.Marshal(mmp)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(mjson))
	fmt.Println(mjson)
}