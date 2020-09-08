package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//url := "http://country.io/capital.json"
	url := "http://httpbin.org/json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)

	w.Write([]byte("hello_server, world"))
}
func main() {
	http.HandleFunc("/", ServeHTTP)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
