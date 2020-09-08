package mylog

import (
	"fmt"
	"os"
	"time"
)

// 往终端打印
type ConsoleLogger struct {
	level Level
}

func NewConsoleLogger(level string) *ConsoleLogger {
	logLevel := parseLogLevel(level)
	c1 := &ConsoleLogger{
		level: logLevel,
	}
	return c1
}

// 将公用的记录日志的功能封闭成一个单独的方法
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	// 不记录大于当前日志等级的日志
	if c.level > level { // f.level是你传进来的一个值,就是New的时候赋的值
		return
	}
	fileName, funcName, line := GetCallerInfo(3)
	// [2019-04-21 15:02:03] [DEBUG] main.go:main.go [14] message:id为10一直在尝试登陆
	logLevel := getLevelStr(level)
	now := time.Now().Format("[ 2006.01.02 15:04:05.000 ]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s",
		now,
		//getLevelStr(level),
		logLevel,
		fileName,
		funcName,
		line,
		format)
	fmt.Fprintf(os.Stdout, format, args...) // 先用args替换 format里面的占位符,然后写到f.logFile里面
	fmt.Fprintln(os.Stdout)
}

// 记录日志
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

// 标准输出不需要关闭
func (c *ConsoleLogger) Close() {
}
