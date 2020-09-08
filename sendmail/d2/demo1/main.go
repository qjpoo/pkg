package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	//e.From = "pingyeaa@163.com"
	e.From = "平也 <qjpoo@163.com>"
	e.To = []string{"120117282@qq.com"}
	e.Subject = "发现惊天大秘密！"
	e.Text = []byte("平也好帅好有智慧哦~")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "qjpoo@163.com", "", "smtp.163.com"))
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
}
