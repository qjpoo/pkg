package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件，块读取 Read()
func readBlock()  {
	fp, err := os.Open("./bufio/main.go")
	if err != nil {
		fmt.Println("open file failure ...")
		return
	}
	defer fp.Close()

	buf := make([]byte, 1024)

	//fmt.Println(fp.Read(buf))

	for {
		//n, err := fp.Read(buf)
		n, err := fp.Read(buf)
		if err == io.EOF {
			fmt.Println("file read finish ...")
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

// 读取文件，行读取 ReadBytes('\n')
func readLine()  {
	fp, err := os.OpenFile("./bufio/1.txt",os.O_RDONLY, 6)
	if err	!= nil {
		fmt.Println("打开文件出错 ...")
		return
	}

	defer fp.Close()

	// 创建文件的缓存区
	r := bufio.NewReader(fp)
	for {
		// 读取文件，行读取 readBytes
		s, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(string(s))
	}

}
func main() {
	// readBlock()
	readLine()

}
