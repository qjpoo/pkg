package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age int
	Addr string
}

func main() {
	// template使用


	http.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		// 打开模版文件
		t, err := template.ParseFiles("./template/demo2/index.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		// 用数据渲染模版
		//data :="go language"
		//t.Execute(writer, data)

		//
		userMap := map[int]Person{
			1: {"chiling", 58, "HongKong"},
			2: {"qujian", 20, "WuHan.China"},
			3: {"xiaoming", 30, "Sz.china"},
		}
		t.Execute(writer, userMap)

	})

	http.HandleFunc("/struct", func(writer http.ResponseWriter, request *http.Request) {
		// 打开模版文件
		t, err := template.ParseFiles("./template/demo2/struct.html")
		// t.Name()  模版的名称就是struct
		if err != nil {
			fmt.Println(err.Error())
		}

		p := Person{"chiling", 48, "Hongkong"}

		// 用数据p渲染模版t
		t.Execute(writer, p)

	})

	http.HandleFunc("/range", func(writer http.ResponseWriter, request *http.Request) {
		//在模版中自定义的方法要在parseFiles模板文件之前添加
		// 自定义函数
		//kuaFunc := func(arg string) (string, error) {
		//	return  arg + " good ...", nil
		//}
		// 2. 把自定义的函数告诉模版
		// tempalte.New() 创建一个template对象
		// template.New("range").Funcs(template.FuncMap{"kua": "kuaFunc"}) 给模版添加自定义的函数
		// 解析模版
		//template.New("range").Funcs(template.FuncMap{"kua": "kuaFunc"}).Parse()  // 追加自定义函数



		// 打开模版文件
		t, err := template.ParseFiles("./template/demo2/range.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		// 用数据渲染模版
		userMap := map[int]Person{
			1: {"chiling", 58, "HongKong"},
			2: {"qujian", 20, "WuHan.China"},
			3: {"xiaoming", 30, "Sz.china"},
		}
		t.Execute(writer, userMap)
	})

	http.ListenAndServe(":8000", nil)



}
