package  main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("get err: ", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
