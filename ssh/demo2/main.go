package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

func main() {
	// 1. 发送指令执行 session.Run()
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
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
