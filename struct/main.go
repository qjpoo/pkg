package main

import "fmt"

type Any *int

// 接收者不能是指针类型
/*func (in Any)log()  {
}*/
func main() {
	a:=[]int{1,2,3,4,5}
	for index,value :=range a{
		fmt.Printf("value:%d valueAddr:%d element:%X\n",value,&value,&a[index])
	}
}
