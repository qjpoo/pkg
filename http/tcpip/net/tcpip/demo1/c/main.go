package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// client

	//conn, err := net.Dial("tcp", "127.0.0.1:8000")
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()

	msg := "GET / HTTP/1.1\r\n"
	//msg += "Host: 127.0.0.1\r\n"
	msg += "Host: www.baidu.com\r\n"
	msg += "Connection: close\r\n"
	msg += "\r\n\r\n"

	// 向conn里面写msg
	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 不断的从conn里面读取数据
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("读取的数据为： ", string(buf[:count]))
	}


}
