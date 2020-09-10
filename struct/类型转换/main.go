package main

import "fmt"

func main() {
	var e interface{}
	e = 10
	switch v := e.(type) {
	case int:
		fmt.Println("整型", v)
		var s int
		s = v
		//fmt.Println(s)
		fmt.Printf("s type: %T, s: %d\n", s, s )
		fmt.Printf("e type: %T, e: %d\n", e, e )

	case string:
		fmt.Println("字符串", v)
	}

	var f interface{} = "xx"
	//ss := f  // 不能这样直接操作,否则会报错
	ss := f.(string)
	fmt.Println(ss)

	bb := fmt.Sprintf("%s", f)
	fmt.Println(bb)

	aa, ok := f.(string)
	if !ok {
		fmt.Println("not string type ...")
		return
	}
	fmt.Println(aa)

	// interface {}转换为[] int
	i1 := []interface{}{1,2,3,4,5}
	fmt.Println("i1: ", len(i1))
	i2 := make([]int,len(i1))
	for i:=0;i<len(i1);i++ {
		i2[i]=i1[i].(int)
	}
	fmt.Println(i2)
}
