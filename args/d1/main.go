package main

import (
	"fmt"
	"os"
)

func main() {
	// os.args

	args := os.Args
	for k, v := range args {
		fmt.Println(k, "===", v)
		/*
		go run main.go -c xxoo 1 2 3
		0 === C:\Users\86137\AppData\Local\Temp\go-build981587782\b001\exe\main.exe
		1 === -c
		2 === xxoo
		3 === 1
		4 === 2
		5 === 3
		*/

		//var xxoo *string
		//str := "A"
		//xxoo = &str
		//fmt.Println(str, xxoo, *xxoo)
	}
}
