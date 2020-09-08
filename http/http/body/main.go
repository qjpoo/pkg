package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//fmt.Printf("resp: %T, resp: %+#v", resp, resp)
	//t := reflect_typeof.TypeOf(resp)
	//v := reflect_typeof.ValueOf(resp)
	//fmt.Println("---------------------")
	//fmt.Println(t.NumField())
	//fmt.Println(t, v)
	/*
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
	}

	 */
	b, err := json.Marshal(resp)
	if err !=nil {
		fmt.Println(err)  // json: unsupported type: func() (io.ReadCloser, error)
	}
	for k, v := range b {
		fmt.Println(k, "------------->", v)
	}
	//fmt.Println(b)


}
