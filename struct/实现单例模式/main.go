package main

import "fmt"

type Single struct {
}

var single *Single

func GetSingle() *Single {
	if single == nil {
		single = &Single{}
	}
	return single
}

func main() {
	fmt.Printf("%p\n", GetSingle())
	fmt.Printf("%p\n", GetSingle())
	fmt.Printf("%p\n", GetSingle())
}
