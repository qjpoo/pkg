package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
)

func main() {
	// 发送指令执行 session.Output(),  session.run(command)是直接在host执行命令，不关心执行结果。session.Output是将执行命令之后的Stdout返回
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", "192.168.11.158:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("dg-mall.com")},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil

		},
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	// 建立新会话
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	result, err := session.Output("date;ifconfig")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to run command, Err:%s", err.Error())
		os.Exit(0)
	}
	fmt.Println(string(result))
}