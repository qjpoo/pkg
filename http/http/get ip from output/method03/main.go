package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)

func main() {
	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

	// You can specify key path by providing arguments to Get function
	v1, v2, v3, v4  := jsonparser.Get(data, "person", "name", "fullName")
	fmt.Println(string(v1), v2,v3,v4)  // Leonid Bugaev string 112 <nil>

	// There is `GetInt` and `GetBoolean` helpers if you exactly know key data type
	v, err := jsonparser.GetInt(data, "person", "github", "followers")
	fmt.Println(v, err)  // 109 <nil>

	// When you try to get object, it will return you []byte slice pointer to data containing it
	// In `company` it will be `{"name": "Acme"}`
	vn,_, _, _ := jsonparser.Get(data, "company", "name")
	fmt.Println(string(vn))

	vx,_, _, _ := jsonparser.Get(data, "person","avatars", "[0]", "url")
	fmt.Println(string(vx)) // https://avatars1.githubusercontent.com/u/14009?v=3&s=460

	// If the key doesn't exist it will throw an error
	var size int64
	if value, err := jsonparser.GetInt(data, "company", "size"); err == nil {
		size = value
	}else {
		fmt.Println(err)
	}
	fmt.Println(size)

	fmt.Println("----------------------------")
	// You can use `ArrayEach` helper to iterate items [item1, item2 .... itemN]
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(jsonparser.Get(value, "url"))
		v, _, _, _ := jsonparser.Get(value, "url")
		fmt.Println(string(v))
	}, "person", "avatars")




}
