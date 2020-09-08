package main

import "fmt"

func main() {
	a := [...]int{2, 3, 1, 0, 2, 5, 3}
	num := make(map[int]bool)
	for _, v := range a {
		if !num[v] {
			fmt.Println(num[v])
			num[v] = true
		} else {
			fmt.Println(v)
		}
	}
}
