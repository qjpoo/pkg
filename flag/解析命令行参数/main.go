package main

import (
	"flag"
	"fmt"
)

func main() {
	var l = flag.Int("l", 0, "para l is for len")
	var c = flag.String("c", "nothing", "para c is for content")
	flag.Parse()
	fmt.Println(*l)
	fmt.Println(*c)
	flag.Usage()
}