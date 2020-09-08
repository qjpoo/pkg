package main

import (
	"pkg/mylog" // 从src开始
)

var logger mylog.Logger

func main() {
	// 在mylogger这个包里面,调用使用自己写的mylog打印日志的那个包

	// 只会打印info及以上的日志
	logger = mylog.NewFileLogger("debug", "./mylogger", "info.log")
	defer logger.Close()
	for {
		userId := 1
		logger.Debug("id是%d的用户一直在尝试登陆", userId)
		logger.Info("id是%d的用户一直在尝试登陆", userId)
		logger.Error("id是%d的用户一直在尝试登陆", userId)
	}
/*
	// 向终端写日志
	logger = mylog.NewConsoleLogger("debug")
	userId := 10
	logger.Debug("id是%d的用户一直在尝试登陆", userId)
	logger.Info("id是%d的用户一直在尝试登陆", userId)
	logger.Error("id是%d的用户一直在尝试登陆", userId)

 */


}

// 按照时间切分, 一个小时一切
// 文件日志的结构体里面记录一下上一次切分的小时
// 每一次写日志的时候,拿当前的时间小时数和上一次切分的小时做对比
// 如果小时数不一致就切分,否则就不切分