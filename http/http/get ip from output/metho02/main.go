package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)

var data = []byte(`{
  "type":"example",
  "data":{
    "name":"abc",
    "labels":{
      "key":"value"
    }
  },
  "subsets":[
    {
      "addresses":[
        {
          "ip":"192.168.103.178"
        }
      ],
      "ports":[
        {
          "port":80
        }
      ]
    }
  ]
}`)

func main() {
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			v, _, _, err := jsonparser.Get(value, "ip")
			if err != nil {
				return
			}
			fmt.Println("ip: ", string(v[:]))
		}, "addresses")
	}, "subsets")

	v, t, _, _ := jsonparser.Get(data, "subsets", "[0]","addresses", "[0]", "ip")
	fmt.Println(string(v), t)
}
