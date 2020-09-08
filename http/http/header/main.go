package main

import (
	"fmt"
	"net/http"
	"time"
)

var url = []string{
	"https://www.baidu.com",
	"https://www.qq.com",
	"https://taobao.com",
	"https://google12wqwqw.com",
}
func main() {
	for _, v := range  url{
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Head(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp.Status)
	}

}
