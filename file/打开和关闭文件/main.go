package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开和关闭文件
	// 只读方式打开当前目录
	f, e := os.Open("./123.log")
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("open sucessfull ...")
	defer f.Close()

}