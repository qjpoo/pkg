package main

import (
	"fmt"
	"sync"
	"time"
)

type Single2 struct {

}

var single2 *Single2
var lock *sync.Mutex = &sync.Mutex{}

func GetSingle2() *Single2 {
	if single2 == nil {
		lock.Lock()
		defer lock.Unlock()
		if single2 == nil {
			single2 = &Single2{}
		}
	}
	return single2
}

func main() {
	go fmt.Printf("%p\n", GetSingle2())
	go fmt.Printf("%p\n", GetSingle2())
	go fmt.Printf("%p\n", GetSingle2())
	time.Sleep(time.Second)
	fmt.Println("done")
}