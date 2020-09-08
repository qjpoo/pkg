//package 声明开头表示代码所属包
package main

import "fmt"

func main() {
	//声明空切片
	var s1 []int

	//:=声明
	s2 :=[]int{}

	var s3 []int = make([]int, 0)
	s4 :=make([]int, 0, 0)
	fmt.Println(s1, s2, s3, s4)
}