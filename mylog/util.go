package mylog

import (
	"path"
	"runtime"
)

// 的得当前运行的文件名和哪一行
func GetCallerInfo(skip int) (fileName, funcName string, line int) {
	pc, fileName, line, ok := runtime.Caller(skip)  // 0是当前, 1是上一层
	if !ok {
		return
	}
	//pc ==> 拿到当前执行的函数名
	funcName = runtime.FuncForPC(pc).Name() // 函数名
	funcName = path.Base(funcName)
	fileName = path.Base(fileName) // 文件名
	//fmt.Println(funName, fileName, line, ok)
	//return fileName, funcName, line
	return
	
}