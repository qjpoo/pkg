package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Test interface{}

type Test1 interface {
	TestFunc()
}

type Structure struct { // Structure实现了Test1接口
	a int
	b string
	c bool
}

func (s *Structure) TestFunc() {
	fmt.Println("Ok, Let's rock and roll!")
}

func fTest(t Test) { // interface
	fmt.Println("t: ", t,t == nil)
}

func fTest1(t1 Test1) { // interface
	fmt.Println("t1: ", t1,t1 == nil)
	fmt.Println(t1 == nil)
}

func fStructure(s *Structure) {
	fmt.Println("s: ", s,s == nil)
	fmt.Println(s == nil)
}

func main() {
	var s *Structure = nil
	fmt.Printf("%T, %v\n", s, s)
	fTest(s)      // false 值为空，类型不为空
	fTest1(s)     // false
	fStructure(s) // true
	s.TestFunc()  // Ok, Let's rock and roll!


	fmt.Println("-------------------------")
	w1 := errors.New("ERROR ...")
	fmt.Printf("%T, %v,%v\n", w1, w1, &w1)   // *errors.errorString, ERROR ...
	s1 := w1.Error()
	fmt.Printf("%T, %v, %v\n", s1, s1, &s1)  // string, ERROR ...

	var p int
	p = 1
	var p1 *int
	p1 = &p
	fmt.Printf("%T, %v, %v\n", p1, p1, *p1)
	fmt.Println(p, p1,&p1)

	fmt.Println("--------------ERR---------------")
	wa := errors.New("ERR")
	fmt.Printf("%T, %v\n", wa, &wa)
	wb := errors.New("ERR")
	fmt.Printf("%T, %v\n", wb, &wb)
	fmt.Println(wa == wb)


	rd, _ := os.Open("123.log")
	r := bufio.NewReader(rd)
	fmt.Printf("%T, %v", r, r)


	xxoo := &Structure{}
	fmt.Printf("%T, %v\n", xxoo, xxoo)
	fmt.Printf("%T, %v\n", xxoo,  *xxoo)
}
