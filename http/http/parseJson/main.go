package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// curl localhost:8000 -d '{"name":"Hello"}'
func Cleaner(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	fmt.Println("1: ", string(b))
	err = json.Unmarshal(b, &msg)  // json ==> struct
	fmt.Println("2: ", msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	fmt.Println("3: ", msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func main() {
	http.HandleFunc("/", Cleaner)
	address := ":8000"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
/*
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var u User
        if r.Body == nil {
            http.Error(w, "Please send a request body", 400)
            return
        }
        err := json.NewDecoder(r.Body).Decode(&u)
        if err != nil {
            http.Error(w, err.Error(), 400)
            return
        }
        fmt.Println(u.Id)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
 */