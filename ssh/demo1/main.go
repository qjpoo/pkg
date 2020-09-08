package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

type ClientConfig struct {
	Host       string
	Port       int64
	Username   string
	Password   string
	Client     *ssh.Client
	LastResult string
}

func (cliConf *ClientConfig) createClient(host string, port int64, username, password string) {
	var (
		client *ssh.Client
		err    error
	)
	cliConf.Host = host
	cliConf.Port = port
	cliConf.Username = username
	cliConf.Password = password

	config := ssh.ClientConfig{
		User: cliConf.Username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	//fmt.Printf("%T, %v\n", ssh.Password(password), ssh .Password(password))

	addr := fmt.Sprintf("%s:%d", cliConf.Host, cliConf.Port)

	// 获取client
	if client, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Fatalln("error occurred: ", err)
	}

	cliConf.Client = client
}

func (cliConf *ClientConfig) RunShell(shell string) string {
	var (
		session *ssh.Session
		err     error
	)

	// 获取session
	if session, err = cliConf.Client.NewSession(); err != nil {
		log.Fatalln("error occured: ", err)
	}

	// 执行shell
	if output, err := session.CombinedOutput(shell); err != nil {
		log.Fatalln("error occured: ", err)
	} else {
		cliConf.LastResult = string(output)
	}
	return cliConf.LastResult

}

func main() {

	cliConf := new(ClientConfig)
	cliConf.createClient("192.168.11.158", 22, "root", "dg-mall.com")

	fmt.Println(cliConf.RunShell("cd /tmp/; ls -al"))

}
