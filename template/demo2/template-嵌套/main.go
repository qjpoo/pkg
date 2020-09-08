package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/demo2/template-嵌套/t.tmpl", "./template/demo2/template-嵌套/ul.tmpl")
	//tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main() {
	// template的嵌套
	//fmt.Println(os.Getwd())
	http.HandleFunc("/tmpl", tmplDemo)


	http.ListenAndServe(":8000", nil)

}
