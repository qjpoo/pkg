package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	// 解析ini问题

	cfg, err := ini.Load("./ini/my.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取操作, 默认分区可以使用空字符串表示
	fmt.Println("App mode: ", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path: ", cfg.Section("paths").Key("data").String())

	//
	fmt.Println("Server Protocol: ", cfg.Section("server").Key("protocol").In("http", []string{"imap", "smtp"}))

	// 如果protocol的值不是imap,或者smtp,就用默认的smtp
	fmt.Println("Email Protocol: ",
		cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))


	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	//保存某些值
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("./ini/hello/my.ini.local")

}
