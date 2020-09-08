package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server ...")
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for{
		conn, err := lis.Accept()
		if err !=nil {
			fmt.Println(err.Error())
			return
		}
		go process(conn)

	}

}

func process(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("read: ", string(buf))
	}

}
