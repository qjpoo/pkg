package mylog

import "strings"

type Level uint16
const (
	DebugLevel Level = iota
	InfoLevel
	ErrorLevel
)

func getLevelStr(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case ErrorLevel:
		return "ERROR"
	default:
		return "INFO"
	}

}


// 根据用户传入的字符串类型的日志级别, 解析也对就的level

func parseLogLevel( levelStr string)  Level {
	// 转化成全小写
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "error":
		return ErrorLevel
	default:
		return DebugLevel

	}

}


// 定义一个logger接口
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	Close()
}