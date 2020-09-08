package main

import (
	"fmt"
	"net/url"
)

func main() {

	var urlStr = "http://baidu.com/index.php/?abc=1_羽毛"
	l, err := url.ParseQuery(urlStr)
	fmt.Println(l, err)

	l2, e2 := url.ParseRequestURI(urlStr)
	fmt.Println(l2, e2)

	l3, e3 := url.Parse(urlStr)
	fmt.Println(l3, e3)
	fmt.Println("path: ",l3.Path, "\n", "RawPath: ",l3.RawPath, "\n", "Query: ", l3.Query(), "\n","query_encode", l3.Query().Encode(), "\n","RequestURI: ", l3.RequestURI())
}
