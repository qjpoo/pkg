package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadFile("./template/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		num := rand.Intn(10)
		dataStr := string(b)
		if num >5 {
			dataStr = strings.Replace(dataStr, "{{data}}", "<li>go语言</li>", 1)
		}else {
			dataStr = strings.Replace(dataStr, "{{data}}", "<li>c语言</li>", 1)
		}
		writer.Write([]byte(dataStr))
	})

	http.ListenAndServe(":8000", nil)
}
