package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {

	url := "https://httpbin.org/json"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	b, e := ioutil.ReadAll(res.Body)
	if e != nil {
		panic(err.Error())
	}

	//fmt.Println(string(b))
	var data interface{} // TopTracks
	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Type: %T, Results: %v\n", data, data)


}
