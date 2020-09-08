package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target *Foo) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	fmt.Println("x", json.NewDecoder(r.Body))
	fmt.Println("xx", json.NewDecoder(r.Body).Decode(target))
	return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
	Bar interface{}
	//Bar string
}

func main() {
	foo1 := new(Foo) // or &Foo{}
	getJson("https://httpbin.org/json", foo1)
	println(foo1.Bar)
}
