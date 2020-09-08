package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// log

	//log.Fatal("xxoo")
	//str := "xxoo"
	//log.Fatalf("11111111111111%s",str)

	//log.Fatalln("222222222222222")
	//fmt.Println("3333333333333")


/*	l := log.New(os.Stdout,"[ Info ] ", 2)
	//fmt.Println(l.Prefix())
	l.Output(2, "hello_server, world ...")
*/

	logger := log.New(os.Stdout, "a.log", log.Lshortfile)
	logger.SetFlags(log.Lshortfile)
	fmt.Println(logger.Flags())

	logger.SetPrefix("[debug]: ")
	fmt.Println(logger.Prefix())

	logger.Println("hello_server, log file")

	err := logger.Output(1, "xo")
	fmt.Println(err)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("xxoo")  // 2020/08/03 11:11:55.216648 C:/Users/86137/go/src/eg/log/main.go:38: xxoo

	// 设置前缀
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[DEBUG]: ")
	log.Println("hello_server...")


	logFile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[DEBUG]: ")
	log.Println("1111111111111111111111111")
	log.Println("2222222222222222222222222")

}
