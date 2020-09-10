package main

import (
	"fmt"
)

func main() {
	var a []int

	if a == nil {
		fmt.Println("it is nil, 1")
	} else {
		fmt.Println("it not not nil, 1")
	}

	fmt.Println("slice len1 is", len(a))

	a = make([]int, 0) // 不必要

	if a == nil {
		fmt.Println("it is nil, 2")
	} else {
		fmt.Println("it not not nil, 2")
	}

	fmt.Println("slice len2 is", len(a))

	a = append(a, 1)

	var m map[int]int

	if m == nil {
		fmt.Println("it is nil, 3")
	} else {
		fmt.Println("it not not nil, 3")
	}

	fmt.Println("map len3 is", len(m))

	m = make(map[int]int, 0) // 必要，否则panic

	if m == nil {
		fmt.Println("it is nil, 4")
	} else {
		fmt.Println("it not not nil, 4")
	}

	fmt.Println("map len4 is", len(m))

	m[1] = 10

	fmt.Println(a, m)

	fmt.Println("---------------->")
	var test = map[string]string{"姓名": "李四", "性别": "男"}
	name, ok := test["姓名1"] // 假如key存在,则name = 李四 ，ok = true,否则，ok = false
	if ok {
		fmt.Println(name)
	} else {
		fmt.Println("没有key", name, "xxx") // 空
	}
	fmt.Println(test["jkjk"]) // 空
	if test["xxoo"] == "" {
		fmt.Println("test[\"xxoo\"]为空")
	}
	delete(test, "姓名") //删除为姓名为key的值，不存在没关系
	fmt.Println(test)
	var ay map[string]string
	//ay["b"] = "c" //这样会报错的，要先初始化内存
	ay = make(map[string]string)
	ay["b"] = "c" //这样才不会错

	fmt.Println("---------------->")
	var m1 map[string]string
	fmt.Println(m1["1"]) // 空
	fmt.Printf("m1[\"1\"]=%#+v\n", m1["1"])

	fmt.Println("---------------->")
	var m2 map[string]string
	fmt.Println(m2) // map[]

	fmt.Println("---------------->")
	var m3 map[string]string
	m3 = make(map[string]string)
	fmt.Println(m3) // map[]

	fmt.Println("---------------->")
	var aa interface{}
	fmt.Println(aa)

	var bb *BB
	fmt.Println(bb)

	aa = bb
	fmt.Println(aa)

	if aa != nil {
		fmt.Println("空接口赋值之后其值为nil, 但是其本身不是nil哟.")
		fmt.Println("所以打印的时候是nil, 判等的时候并不是nil.")
		fmt.Println(aa)
	}
	fmt.Printf("%T, %#v\n", aa, aa)  // *main.BB, (*main.BB)(nil)


	fmt.Println("---------------->")
	// 空的map
	my_map := map[string]string{}
	fmt.Println(my_map)  // map[]

	// nil map
	var my_map_1 map[string]string
	fmt.Println(my_map_1)  // map[]




}

type BB struct {
}
