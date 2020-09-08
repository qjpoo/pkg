package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// 打印出cookie
	cookies := r.Cookies()
	//fmt.Fprint(w, cookies)
	log.Println(cookies)
	r.ParseForm()
	// 解析参数
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{
		Name:    "golang",
		Value:   "7777",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	t, _ := template.ParseFiles("./cookie/储存cookie/index.html")
	t.Execute(w, t)
	// 这样就可以直接返回一个index.html首页
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.ListenAndServe("0.0.0.0:8000", nil)

}
