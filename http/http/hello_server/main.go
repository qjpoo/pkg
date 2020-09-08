package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("http listen failure ...")
	}
}

func Hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle hello_server")
	//fmt.Fprintln(w, "hello_server")
	io.WriteString(w, "xxoo")


}
