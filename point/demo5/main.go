package main

import "fmt"

func change(x *int)  {
	*x = 60
}
func main() {

	a := 58
	fmt.Println(" before a: ", a)
	b := &a
	change(b)
	fmt.Println(" after a: ", a)


	arr := [3]int{1,2,3}
	modify(&arr)
	fmt.Println(arr)

	a1 := []int{1,2,3}
	modify1(a1)
	fmt.Println("a: ", a1)

	var i int
	var ptr *int
	var pptr **int
	i = 100
	ptr = &i
	pptr = &ptr

	fmt.Println(i,&i, ptr, pptr)
	fmt.Println(*pptr, &ptr, *pptr, &pptr)

}
func modify1(x []int)  {
	x[0] = 200


}

func modify(x *[3]int)  {
	(*x)[0] = 100
}