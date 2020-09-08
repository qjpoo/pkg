package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


/*
{
  "success": true,
  "message": "''",
  "result": [
    {
      "High": 0.0135,
      "Low": 0.012,
      "Created": "2014-02-13T00:00:00"
    }
  ]
}
 */

type Summary struct {
	Sucess bool `json:"success"`
	Message string `json:"message"`
	Result []Result `json:"result"`
}

type Result struct {
	Created string  `json:"Created"`
	High    float64 `json:"High"`
	Low     float64 `json:"Low"`
}

func getSummary() {

	url := "http://myurl"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err.Error())
	}

	log.Printf("body = %v", string(body))
	//outputs: {"success":true,"message":"","result":["High":0.43600000,"Low":0.43003737],"Created":"2017-06-25T03:06:46.83"}]}

	var summary = new(Summary)
	err3 := json.Unmarshal(body, &summary)
	if err3 != nil {
		fmt.Println("whoops:", err3)
		//outputs: whoops: <nil>
	}

	log.Printf("s = %v", summary)
	//outputs: s = &{{0 0 0  0 0 0  0 0 0  0}}

}

func main()  {
	getSummary()
}