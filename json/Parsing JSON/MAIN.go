package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Number int `json:"number"`
	Message string `json:"message"`
	People []People `json:"people"`
}

type People struct {
	Craft string `json:"craft"`
	Name string `json:"name"`
}

func main() {
	text := `{
"people": 
[
{"craft": "ISS", "name": "Sergey Rizhikov"}, 
{"craft": "ISS", "name": "Andrey Borisenko"}, 
{"craft": "ISS", "name": "Shane Kimbrough"}, 
{"craft": "ISS", "name": "Oleg Novitskiy"}, 
{"craft": "ISS", "name": "Thomas Pesquet"}, 
{"craft": "ISS", "name": "Peggy Whitson"}
], 
"message": "success", 
"number": 6
}`
	textBytes := []byte(text)

	data := Data{}
	err := json.Unmarshal(textBytes, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data.Message)
	fmt.Println(data.People[0])
	fmt.Println(data.Number)
}
