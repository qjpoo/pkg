package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// filepath

	// 当看当前系统的分隔符
	sep := os.PathSeparator
	fmt.Println(string(sep)) //   \

	// 将分割符转为 /
	// 注意：该函数不在意路径是否存在，只是当成普通的字符串进行处理
	fmt.Println(filepath.ToSlash(`c:\windows\system\123.exe`)) // c:/windows/system/123.exe

	// 获取path中最后一个分割符前面的部分，类似于python中的os.path.dirname(path)
	fmt.Println(filepath.Dir(`c:\windows\system32\123.exe`)) // c:\windows\system32

	// 获取path中最后一个分割符后面的部分
	fmt.Println(filepath.Base(`c:\windows\system32\123.exe`)) // 123.exe

	// split(path string) (dir, file string) 相当于是Dir和Base的组合
	dir, file := filepath.Split(`c:\windows\123.exe`)
	fmt.Println(dir, file) // c:\windows\ 123.exe

	// Ext(path string)
	ext := filepath.Ext(`c:\windows\123.exe`)
	fmt.Println(ext) // .exe

	//dirPath, _ := os.Getwd()

}
