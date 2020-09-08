package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%T\n", reader)
	fmt.Print("Please input something: ")
	text, _ := reader.ReadBytes('\n')
	fmt.Printf("%T\n", text) // []uint8
	fmt.Println(string(text))

	//text, _ := reader.ReadString('\n')
	//fmt.Printf("%T\n", text) //  string
	//fmt.Println(text)

	s := []uint8{100,101,110}
	fmt.Println(string(s))
	fmt.Printf("%v, %c", s, s)

}
