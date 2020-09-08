package main

import "fmt"

//传入数组，修改元素
func printArr(arr [5]int)  {
	//修改一个元素
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}
}

func main() {
	//定义数组
	var arr1 [5]int

	//:=声明并赋值
	arr2 := [3]int{1, 2, 3}
	//可以省略大小
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	//printArr(arr1)
	//报错
	//printArr(arr2)
	printArr(arr3)

	//打印原始值
	fmt.Println()
	fmt.Println(arr3)
}