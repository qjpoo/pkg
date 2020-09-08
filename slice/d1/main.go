package main

import "fmt"

func main() {
	//var slice2 []int=[]int{1,2,3}

	//好像是需要这样才行
	//var a []int
	// 切片名称是a,切片数据类型是int类型

	//arr:=[...]int{1,2,3,4,5}
	//定义的数组

	//a=arr[1:3]
	//赋值给切片
	//注意：数组创建的切片会互相干扰
	/*	var a []int
		// 切片名称是a,切片数据类型是int类型
		arr := [...]int{1, 2, 3, 4, 5}
		a = arr[1:3]

		fmt.Println(a)
		arr[1] = 10
		fmt.Println(a)
		a[1] = 100
		fmt.Println(arr)*/
	/*
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println(a)
		a = append(a[:2], a[2+2:]...)
		// 他是这样删除数据的，把两个新的拼起来。这个例子就是删除了第二个数据
		fmt.Println(a)*/

/*	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 10)
	fmt.Println(b)
	copy(b, a)
	fmt.Println(b)
	// 必须要先make出来，才能复制进去*/



	// 开头添加元素
	var b []int=[]int{1,2,3}
	b=append([]int{0},b...)
	// 开头添加一个0
	fmt.Println(b)

}
