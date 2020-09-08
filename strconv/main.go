package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv

	// int ==> string
	num := 100
	fmt.Println(string(num))

	// itoa   int ==> string
	// 如果想把int转成string应该是用Itoa， i就是integer，这个a是什么？其实这个a就是字符串，
	//只不过这是从C语言中遗留下来的，因为C中没有字符串的概念，而是使用字符数组，
	//所以这个a在C中是array。但是在go中就把a当中字符串即可
	num =100
	fmt.Println(strconv.Itoa(num))
	fmt.Println(strconv.Itoa(num) == "100")

	// atoi (表示 ascii to integer)是把字符串转换成整型数的一个函数
	// 整型转字符串可以直接转，但是字符串转整型，则是有可能发生错误的，因为必须要求字符串的每一个字符必须是数字也可以转
	// 所以这里除了返回结果，还会返回一个error
	ascii := "97"
	fmt.Println(strconv.Atoi(ascii)) // 97 nil
	//str := "98xx"
	// fmt.Println(strconv.Atoi(str)) // 0 strconv.Atoi: parsing "98xx": invalid syntax


	// Parse类函数用于转换字符串为给定类型的值：ParseBool、ParseFloat、ParseInt、ParseUint
	// ParseBool
	// 将指定字符串转换为对应的bool类型，只接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，否则返回错误
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("F"))

	// ParseInt
	/*
	func ParseInt(s string, base int, bitSize int) (i int64, err error)

	s：转成int的字符串
	base：指定进制(2到36)，如果base为0，那么会从字符串的前置来判断，如0x表示16进制等等，如果前缀也没有那么默认是10进制
	bistSize：整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64
	 */
	fmt.Println(strconv.ParseInt("100", 10, 64))

}
