package main

import "fmt"

type person struct {
	name string
}

/*
// This is working, but a really Not how I want to do this
func personConstructor(params ...string) person {
	name := "Unnamed Person" // Default name
	if len(params) > 0 {
		name = params[0]
	}
	return person{name: name}
}

func main() {
	bob := personConstructor("Bob")
	fmt.Println(bob)
}
*/


func (p person) constructor(params ...string) person {
	name := "Unnamed Person" // Default name
	if len(params) > 0 {
		name = params[0]
	}
	return person{name: name}
}

func main() {
	// Lets define a new instance of a person struct
	p := person{}

	// The function, when called on a instance, works as expected
	bob := p.constructor("Bob")

	fmt.Printf("name: %+v", bob)
}