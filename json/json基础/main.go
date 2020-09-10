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

	ret := TraverseJsonToTree(a1)
	fmt.Println(string(ret))

}

func TraverseJsonToTree(json []byte) []byte {
	stringFlag := false
	jsonTree := make([]byte, 0)
	tabNum := 0
	dirFlag := false
	jsonLen := len(json)

	for i, jsChar := range json {
		if stringFlag {
			jsonTree = append(jsonTree, jsChar)
			if jsChar == '"' {
				stringFlag = false
			}
			continue
		}

		if jsChar == '"' {
			jsonTree = append(jsonTree, jsChar)
			stringFlag = true
			continue
		}

		if jsChar == '{' || jsChar == '[' {
			jsonTree = append(jsonTree, jsChar)
			jsonTree = append(jsonTree, '\r')
			jsonTree = append(jsonTree, '\n')

			tabNum++
			for tabIndex := 0; tabIndex < tabNum; tabIndex++ {
				jsonTree = append(jsonTree, '\t')
			}
			continue
		}

		if jsChar == ',' {
			jsonTree = append(jsonTree, jsChar)
			jsonTree = append(jsonTree, '\r')
			jsonTree = append(jsonTree, '\n')

			for tabIndex := 0; tabIndex < tabNum; tabIndex++ {
				jsonTree = append(jsonTree, '\t')
			}
			continue
		}

		if dirFlag {
			dirFlag = false
			tabNum--
			jsonTree = append(jsonTree, '\r')
			jsonTree = append(jsonTree, '\n')

			for tabIndex := 0; tabIndex < tabNum; tabIndex++ {
				jsonTree = append(jsonTree, '\t')
			}

			jsonTree = append(jsonTree, jsChar)
			continue
		}

		if jsChar == ']' || jsChar == '}' {
			tabNum--
			jsonTree = append(jsonTree, '\r')
			jsonTree = append(jsonTree, '\n')

			for tabIndex := 0; tabIndex < tabNum; tabIndex++ {
				jsonTree = append(jsonTree, '\t')
			}

			jsonTree = append(jsonTree, jsChar)

			if (i + 1) < jsonLen {
				if json[i+1] != ',' {
					dirFlag = true
				}
			}
			continue
		}

		if jsChar == ':' || (jsChar >= '0' && jsChar <= '9') {
			jsonTree = append(jsonTree, jsChar)
			continue
		}

	}

	return jsonTree
}
