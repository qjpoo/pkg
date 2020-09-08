package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age string
}

func main() {
	t, err := template.ParseFiles("./template/hello_server/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	p := Person{"chiling", "50"}
	if err := t.Execute(os.Stdout, p); err != nil { // 使用模版t,用p来填充,输出到stdout,html中的.代表p这个对象
		fmt.Println(err)
	}

}
