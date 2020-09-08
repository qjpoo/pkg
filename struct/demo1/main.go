package main

import "fmt"

type Books struct {
	title string
}


func main() {

	var Book1 Books
	var Book2 Books

	Book1.title = "go language"
	Book2.title = "shell scripts"

	fmt.Println(Book1)  // {go language}
	fmt.Println(&Book1)   //  &{go language}
	printBook(&Book1)

	fmt.Println("----------------")
	printBook1(Book2)

}

func printBook( book *Books)  {
	fmt.Println(book, *book)  // &{go language} {go language}
	book.title = "golang ..."
	fmt.Println(book.title)
}

func printBook1( book Books)  {
	fmt.Println(book)  //
	book.title = "java"
	fmt.Println(book.title)
}

