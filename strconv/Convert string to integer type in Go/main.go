package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	s := flag.Arg(0)
	fmt.Println("one: ", s)
	// string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(s, i)
}