package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	i := []int{3, 7, 1, 3, 6, 9, 4, 1, 8, 5, 2, 0}
	a := sort.IntSlice(i) // 转变成IntSlice类型
	fmt.Printf("%T, %v\n", a, a)
	fmt.Println(sort.IsSorted(a)) // false
	sort.Sort(a)
	fmt.Println(a)                      // [0 1 1 2 3 3 4 5 6 7 8 9]
	fmt.Println(sort.IsSorted(a), "\n") // true

	b := sort.IntSlice{3}
	fmt.Println(sort.IsSorted(b), "\n") // true

	// 更改排序行为
	c := sort.Reverse(a)
	fmt.Println(sort.IsSorted(c)) // false
	fmt.Println(reflect.TypeOf(c))  // *sort.reverse
	fmt.Println(c)                // &{[0 1 1 2 3 3 4 5 6 7 8 9]}
	sort.Sort(c)
	fmt.Println(c)                      // &{[9 8 7 6 5 4 3 3 2 1 1 0]}
	fmt.Println(sort.IsSorted(c), "\n") // true

	// 再次更改排序行为
	d := sort.Reverse(c)
	fmt.Println(sort.IsSorted(d)) // false
	sort.Sort(d)
	fmt.Println(d)                // &{0xc42000a3b0}
	fmt.Println(sort.IsSorted(d)) // true
	fmt.Println(d)                // &{0xc42000a3b0}

	fmt.Println("------------------")
	s := []int{5, 2, 3, 0,7, 88, 69}
	inter  := sort.IntSlice(s)
	fmt.Println(reflect.TypeOf(inter), inter)
	sInter := sort.Reverse(inter)
	fmt.Println(reflect.TypeOf(sInter), sInter)
	sort.Sort(sInter)
	fmt.Println(sInter)

	fmt.Println("------------------")
	s = []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

	fmt.Println("------------------")
	s = []int{5, 2, 6, 3, 1, 4} // unsorted
	//sort.Sort(s)
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)

	sort.Sort(sort.IntSlice{1,10,9,8,0,3})
}
