package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// flag.Typevar, 需要先申明变量
	/*
		var user string
		flag.StringVar(&user, "u", "qujian", "用户名，默认为qujian")
		flag.Parse() // 解析上面的定义的标签
		fmt.Println(user) //  go run main.go -u xxoo       xxoo
	*/

	var user, pwd, host string
	var port int
	flag.StringVar(&user, "u", "", "用户名， 默认为空")
	flag.StringVar(&pwd, "p", "", "密码， 默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "P", 3306, "端口号，默认为3306")

	// 从args中解析注册flag
	flag.Parse()

	fmt.Printf("user: %s\npwd: %s\nhost: %s\nport: %d\n", user, pwd, host, port)

	fmt.Println("----------------------------------")
	fmt.Println("user:", flag.Arg(0 ))
	fmt.Println("pwd:", flag.Arg(1))
	fmt.Println("host:", flag.Arg(2))
	fmt.Println("String args:", user)
	fmt.Println("os.Args[0]:", os.Args[0])
	fmt.Println("os.Args[1]:", os.Args[1])
	fmt.Println("os.Args[2]:", os.Args[2])

	fmt.Scan()
}
