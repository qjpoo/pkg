//package 声明开头表示代码所属包
package main

import "fmt"

func main() {
	p := new([]int) // 不会初始化内存
	fmt.Printf("p new: %v\n", p)

	//[]int切片
	//10: 初始化10个长度
	//50: 容量为50
	m := make([]int, 10, 50)
	//m := make([]int,10)
	fmt.Printf("m make: %v\n", m)

	m[0] = 10
	//(*p)[0] = 10
	fmt.Println(p)

	arr1 :=[4]int{1,2,3}
	fmt.Println(arr1)
}
