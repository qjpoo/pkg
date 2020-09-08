package main

import (
	"fmt"
	"net"
)

func main() {
	// udp client


	conn, err := net.Dial("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello_server, world"))
	if err != nil {
		fmt.Println(err.Error())
	}

	// 收消息
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("接收到回复 ", string(buf))

}
