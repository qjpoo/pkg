package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// io
	
	r := strings.NewReader("hello_server, world\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(err)
	}
	// Writestring
	io.WriteString(os.Stdout, "xxoo ...\n")

	// 从标准输入COPY到标准输出
	if n, err := io.Copy(os.Stdout, os.Stdin);err != nil {
		fmt.Println(n, err)
	}
}
