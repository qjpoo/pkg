package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type people struct {
	Number int `json:"number"`
}

func main() {
	// ioutil.ReadAll
	resp, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		fmt.Println(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(body))
	// {"number": 3, "people": [{"craft": "ISS", "name": "Chris Cassidy"}, {"craft": "ISS", "name": "Anatoly Ivanishin"}, {"craft": "ISS", "name": "Ivan Vagner"}], "message": "success"}
	//fmt.Println(body)

	people1 := people{}
	// 把json转为结构体
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	fmt.Println(people1.Number)


}
