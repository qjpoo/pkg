package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./template/demo2/fun/fun.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 自定义一个夸人的模板函数
	kua := func(arg string) (string, error) {
		return arg + " 真帅 !!! ", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	// 相当于创建了一个叫hello的模版
	tmpl, err := template.New("hello_server").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := map[int]UserInfo{
		1: {"chiling", "女", 48},
		2: {"qujian", "男", 28},
		3: {"xiaoming", "男", 18},
	}


	// 使用user渲染模板，并将结果写入w
	tmpl.Execute(w, user)

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
