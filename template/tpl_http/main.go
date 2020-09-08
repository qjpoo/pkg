package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  string
}
var myTmpl *template.Template

func InitTmpl() (err error){
	myTmpl, err = template.ParseFiles("./template/tpl_http/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func main() {
	InitTmpl()
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("http listen failure ...")
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	p := Person{"chiling", "50"}
	if err := myTmpl.Execute(w, p); err != nil {
		fmt.Println(err)
		return
	}
	// 写在文件里
	file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_APPEND, 0666)
	if err !=nil {
		fmt.Println(err)
	}
	if err := myTmpl.Execute(file, p); err !=nil {
		fmt.Println(err)
	}

	// 写到自定义的结构体中,这个结构体要实现Write方法
	res := &Result{}
	myTmpl.Execute(res, p)
	fmt.Println(res.Output)

}

type Result struct {
	Output string
}

func (r *Result) Write(b []byte) (n int, err error)  {
	fmt.Println("result struct implement write interface ...")
	r.Output += string(b)
	return len(b), nil

}
