package main

import "fmt"

type Phone interface {
	call()
}

type Nokia struct {
}

type Iphone struct {
}

func (n Nokia) call()  {
	fmt.Println("nokia ....")
}

func (i *Iphone) call()  {
	fmt.Println( "iphone ...")
}

func main() {
	var phone Phone
	phone = Nokia{}
	phone.call()

	phone = &Nokia{}
	phone.call()

	phone = &Iphone{}
	phone.call()

	i := &Iphone{}
	(*i).call()
	i.call()





}
