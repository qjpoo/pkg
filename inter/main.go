package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// interface

	var w io.Writer  // type Writer interface
	fmt.Printf("w: %T, w: %v\n", w, w) // w: <nil>, w: <nil>

	var pr io.Reader // type Reader interface
	fmt.Printf("pr: %T, pr: %v\n", pr, pr) // pr: <nil>, pr: <nil>

	var s bufio.Scanner  // type Scanner struct
	fmt.Printf("s: %T, s: %v\n", s, s)

	var r bufio.Reader  // type Reader struct
	fmt.Printf("r: %T, r: %v\n", r, r)



	fmt.Println("---------new(io.Writer)-----------")
	var w1 io.Writer  // type Writer interface
	fmt.Printf("w1 before: %T, w1 before : %v\n", w1, w1)  // w1 before: <nil>, w1 before : <nil>
	w2 := new(io.Writer)
	fmt.Printf("w2: %T, w2: %v, w2: %v\n", w2, w2, *w2)  // w2: *io.Writer, w2: 0xc00003a2c0  w2: <nil>
	i := new(int)
	fmt.Printf("i: %T, i: %v\n", i, i)  //  i: *int, i: 0xc0000160b8


	fmt.Println("--------new(bytes.Buffer)------------")
	g := new(bytes.Buffer)
	fmt.Printf("%T, %v\n", g, g)
	n, _ := g.WriteTo(os.Stdout)
	fmt.Println(n)

	fmt.Println("--------------------")
	var k bytes.Buffer
	fmt.Printf("%T, %v\n", k, k)
	k.WriteString("xxoo")

}
