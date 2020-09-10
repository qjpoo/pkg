package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	sli := []byte("abc")
	str := hex.EncodeToString(sli)
	fmt.Printf("%T\n", sli) // []uint8
	fmt.Printf("%T\n", str) // string
	fmt.Println(sli)  // [97 98 99]
	fmt.Println(str)  // 616263

	sli2, _ := hex.DecodeString(str)
	fmt.Printf("%T\n", sli2) // []uint8
	fmt.Println(sli2) // [97 98 99]

	str2 := fmt.Sprintf("%s", sli2)
	fmt.Printf("%T\n", str2)
	fmt.Println(str2)

	fmt.Println(string(sli2))
	fmt.Println([]byte(str2))
}
