package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type User struct {
	Id      string
	Balance uint64
}

func main() {

	u := User{Id: "110", Balance: 8}
	b := new(bytes.Buffer)
	//fmt.Println("x: ", json.NewEncoder(b))
	// 先要生成一个Encoder == >encode解码
	//fmt.Println("o: ", json.NewEncoder(b).Encode(u))
	json.NewEncoder(b).Encode(u)
	res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
	/*
	{
	  "args": {},
	  "data": "{\"Id\":\"110\",\"Balance\":8}\n",
	  "files": {},
	  "form": {},
	  "headers": {
	    "Accept-Encoding": "gzip",
	    "Content-Length": "25",
	    "Content-Type": "application/json; charset=utf-8",
	    "Host": "httpbin.org",
	    "User-Agent": "Go-http-client/2.0",
	    "X-Amzn-Trace-Id": "Root=1-5f432fe7-6a098758728246daf82da30f"
	  },
	  "json": {
	    "Balance": 8,
	    "Id": "110"
	  },
	  "origin": "103.136.251.98",
	  "url": "https://httpbin.org/post"
	}

	*/

	/*
		defer res.Body.Close()
		body,_ := ioutil.ReadAll(res.Body)

		data := make([]data, 0)
		e := json.Unmarshal([]byte(body), &data )
		if e !=nil {
			fmt.Println("error: ", e)
		}
		fmt.Println(data)


	*/

}

/*
type data struct {
	Id      string `json:"ID"`
	Balance int    `json:"Balance"`
}

 */
