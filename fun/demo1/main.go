package main

import "fmt"
import "time"

func goFunc1(f func()) {
	go f()
}

func goFunc2(f func(interface{}), i interface{}) {
	go f(i)
}

func goFunc(f interface{}, args ...interface{}) {
	fmt.Println("args: ", args)
	if len(args) > 1 {
		go f.(func(...interface{}))(args)
	} else if len(args) == 1 {
		go f.(func(interface{}))(args[0])
	} else {
		go f.(func())()
	}
}

func f1() {
	fmt.Println("f1 done")
}

func f2(i interface{}) {
	fmt.Println("f2 done", i)
}

func f3(args ...interface{}) {
	fmt.Println("f3 done", args)
}

func main() {
	goFunc1(f1)
	goFunc2(f2, 100)

	goFunc(f1)
	goFunc(f2, "xxxx")
	goFunc(f3, "hello_server", "world", 1, 3.14)
	time.Sleep(5 * time.Second)
}
