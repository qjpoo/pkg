package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("start server ...")
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err.Error())
	}


	for {
		conn, err := lis.Accept()
		if err != nil {
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
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("read: ", string(buf[:n]))
		fmt.Println("已收到， 请回复")
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		fmt.Scanln(&str)
		fmt.Println(str)
	}
}
