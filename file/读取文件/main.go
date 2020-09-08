package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 基本使用

	f, e := os.Open("./123.log")
	if e != nil {
		fmt.Println(e)
		return
	}
	defer f.Close()

	tmp:= make([]byte, 128)
	n, err := f.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("读取的字节数为: ", n)
	fmt.Println(string(tmp[:n]))


}
