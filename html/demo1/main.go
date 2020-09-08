package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {


	// post
	http.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadFile("./html/form.html")
		if err != nil {
			fmt.Println(err.Error())
		}
		writer.Write(b)
	})

	// register
	http.HandleFunc("/reg", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write([]byte("/reg"))
		fmt.Printf("host: %s, method: %s, form: %s\n", request.Host, request.Method, request.Form)
		//获取注册的信息
		request.ParseForm() // 解析
		//request.FormValue("username")
		//fmt.Printf("request.Form: %#v\n", request.Form)  // request.Form: url.Values{"addr":[]string{"wh"}, "avatar":[]string{""}, "gender":[]string{"male"}, "info":[]string{""}, "		password":[]string{""}, "username":[]string{"guest"}}
		//fmt.Printf("request.PostForm: %#v\n", request.PostForm)  // request.PostForm: url.Values{"addr":[]string{"wh"}, "avatar":[]string{""}, "gender":[]string{"male"}, "info":[]string{""}, "password":[]string{""}, "username":[]string{"guest"}}
		username := request.Form.Get("username") // 这里的key是根据form表单的中Input标签的name属性来决定的
		password := request.Form.Get("password")
		writer.Write([]byte(username + password))





	})


	http.ListenAndServe(":8000", nil)
}
