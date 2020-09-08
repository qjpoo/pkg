package main

import (
	"fmt"
	"net"
)

func main() {
	// udp server

	//udpaddr, err := net.ResolveIPAddr("udp", "127.0.0.0:8000")
	// udpConn, err := net.ListenUDP("udp", udpaddr)

	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8000,
	})
	// 可以通过ParseIp把一个字符串的IP转化成IP类型
	//net.ParseIP("127.0.0.1")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// udp没有对客户端连接的Accept函数,因为它是面向无连接的

	process(udpConn)

}

// 可以接收多个client
func process(udpConn *net.UDPConn) {
	defer udpConn.Close()
	for {
		data := make([]byte, 1024)
		// addr就是client对方的地址
		n, addr, err := udpConn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("接受来自: %v 的消息为: %v, 数量为: %v\n", addr, string(data[:n]), n)

		// 发送数据
		//_, err = udpConn.WriteToUDP(data[:n], addr)
		_, err = udpConn.WriteToUDP([]byte("收到 ..."), addr)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

	}

}
