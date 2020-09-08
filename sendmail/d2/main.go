// main.go
package main

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"net/smtp"
	"time"
)

func SendMail(fromUser, toUser, subject string) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = fromUser
	// 收件人(可以有多个)
	e.To = []string{toUser}
	// 邮件主题
	e.Subject = subject
	// 解析html模板
	t, err := template.ParseFiles("./sendmail/d2/email-template.html")
	if err != nil {
		return err
	}
	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	t.Execute(body, struct {
		FromUserName string
		ToUserName   string
		TimeDate     string
		Message      string
	}{
		FromUserName: "szops@dg-mall.com",
		ToUserName:   "qujian@51-dg.com",
		TimeDate:     time.Now().Format("2006/01/02"),
		Message:      "golang是世界上最好的语言！",
	})
	// html形式的消息
	e.HTML = body.Bytes()
	// 从缓冲中将内容作为附件到邮件中
	e.Attach(body, "email-template.html", "text/html")
	//以路径将文件作为附件添加到邮件中
	e.AttachFile("./sendmail/d2/main.go")
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	fmt.Println("xxxoo")
	return e.Send("smtp.139.com:465", smtp.PlainAuth("", "lovechiling@139.com", "Aa123456", "smtp.139.com"))
}

func main() {
	fromUser := "lovechiling@139.com"
	//toUser := "qujian@51-dg.com, szops@dg-mall.com"
	toUser := "qujian@51-dg.com"
	subject := "hello_server,world"
	err := SendMail(fromUser, toUser, subject)
	fmt.Println(err)
	if err != nil {
		log.Println("发送邮件失败")
		return
	}
	log.Println("发送邮件成功")
}
