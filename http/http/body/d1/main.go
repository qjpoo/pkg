package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Foo struct {
	Bar string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

//func getJson(url string, target interface{}) error {
func getJson(url string, target interface{}) {
	r, err := myClient.Get(url)
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer r.Body.Close()

	//return json.NewDecoder(r.Body).Decode(target)
	//json.NewDecoder(r.Body).Decode(target)
	//v := json.NewDecoder(r.Body).Decode(target)
	//b := make([]byte, 0)
	//fmt.Println(r.Body.Read(b))
	//v := json.NewDecoder(r.Body)
	//fmt.Println(v)
	//fmt.Println(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
}

func main() {
	foo1 := new(Foo) // or &Foo{}
	//v := getJson("https://www.baidu.com", foo1)
	getJson("https://www.baidu.com", foo1)
	//getJson("http://api.open-notify.org/astros.json", foo1)
	//fmt.Println(foo1.Bar)
	//fmt.Println(v)

	// alternately:

	/*
		foo2 := Foo{}
		getJson("http://example.com", &foo2)
		println(foo2.Bar)

	*/
}
