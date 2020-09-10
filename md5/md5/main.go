package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	str := "abc"
	x := md5.New()
	x.Write([]byte(str))
	y := x.Sum([]byte(""))

	fmt.Printf("%x\n", y) // 900150983cd24fb0d6963f7d28e17f72
	fmt.Printf("%T\n", x)  // *md5.digest
	fmt.Printf("%T\n", y)  // []uint8

	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))

	f, err := os.Open("123.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h1 := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", h1.Sum(nil))

	fmt.Println("--------------->")
	j1 := md5.Sum([]byte("abc123"))
	fmt.Println(j1)
	fmt.Printf("%x", j1)
}
