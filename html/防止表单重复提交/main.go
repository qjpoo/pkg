package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/防止表单重复提交/login.html")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":8000", nil) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
