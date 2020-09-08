package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	// 循环读取
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./123.log")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()

	// 循环读取文件
	var b []byte
	var tmp = make([]byte, 128)
	//var tmp  []byte
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		b = append(b, tmp[:n]...)
	}

}