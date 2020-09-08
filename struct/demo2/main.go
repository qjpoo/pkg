package main

import "fmt"

func printNumber(ch chan struct{})  {
	for i:=1;i<=10;i++{
		fmt.Println("printNumber func: ", i)
	}
	data := <- ch
	fmt.Println("data: ", data)

}
func main() {
	ch := make(chan struct{})
	go printNumber(ch)

	ch <- struct{}{}


	fmt.Println("main end ....")

}
