package mylog

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// FileLogger 往文件中记录日志的结构
type FileLogger struct {
	//level int
	level       Level
	logFilePath string
	logFileName string
	logFile     *os.File // os包中File类型指针
	errFile     *os.File
	maxSize     int64
}

func NewFileLogger(level, logFilePath, logFilename string) *FileLogger {
	logLevel := parseLogLevel(level)
	f1Obj := &FileLogger{
		level:       logLevel,
		logFilePath: logFilePath,
		logFileName: logFilename,
		maxSize:     10 * 1024 * 1024, // 10M
	}
	f1Obj.initFileLogger() // 调用下面的初始化方法
	return f1Obj
}

// 专门用来初始化文件日志的文件句柄
func (f *FileLogger) initFileLogger() {
	// 拼接一个filePath字符串
	//filePath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	logName := path.Join(f.logFilePath, f.logFileName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file: %s failed ...", logName))
	}

	f.logFile = fileObj

	// 打开错误日志的文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败", errLogName))
	}
	f.errFile = errFileObj

}

// 检查是否需要拆分
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fileSize := fileInfo.Size()
	// 当传进来的日志大于设置的最大文件时候,切分
	return fileSize >= f.maxSize
}

// 检查按小时切割是否需要拆分
func (f *FileLogger) checkSplitByHour(file *os.File) bool {
	// 获取当前file的时间
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	tmpTimeStr := stat.ModTime().Format("2006.01.02 15:04:05")

	//fmt.Println(tmpTimeStr)
	//fmt.Println(stat.ModTime().Format("2006.01.02 15:04:05"))
	tmpHourStr := strings.Split(tmpTimeStr, " ")
	h := tmpHourStr[1] //15:04:05
	hourStr := strings.Split(h, ":")
	//fmt.Println(hourStr[0]) // 打印文件修改的时间
	tmpCurHour := time.Now().Hour() // int
	curHour := strconv.Itoa(tmpCurHour) // string

	if hourStr[0] != curHour {
		// 不相等就要切
		return true
	} else {
		return false
	}

}
func (f *FileLogger) splitLogFileByHour(file *os.File) *os.File {
	// 得到当前时间的小时
	curTime := time.Now().Format("2006.01.02 15:04:05")
	tmpCurHour := strings.Split(curTime, ":")
	curHour := tmpCurHour[1]
	//如果不相等,就要切割
	fileName := file.Name() // 拿到文件的完整路径
	fmt.Println(fileName)   // mylogger/info.log

	backupName := fmt.Sprintf("%s_%v.bak_%v", fileName, time.Now().Unix(), curHour)
	// 1. 把原来的文件关闭
	file.Close()
	// 2. 备份原来的文件
	os.Rename(fileName, backupName)
	// 3. 新建一个
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	f.logFile = fileObj
	return fileObj
}


// 封闭一个切分日志的文件的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	// 切分
	fileName := file.Name() // 拿到文件的完整路径
	fmt.Println(fileName)   // mylogger/info.log

	backupName := fmt.Sprintf("%s_%v.bak", fileName, time.Now().Unix())
	// 1. 把原来的文件关闭
	file.Close()
	// 2. 备份原来的文件
	os.Rename(fileName, backupName)
	// 3. 新建一个
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	f.logFile = fileObj
	return fileObj
}

// 将公用的记录日志的功能封闭成一个单独的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	// 不记录大于当前日志等级的日志
	if f.level > level { // f.level是你传进来的一个值,就是New的时候赋的值
		return
	}
	fileName, funcName, line := GetCallerInfo(2)
	// 往文件里面写
	// 往哪个日志文件里面写
	// 日志的格式要丰富:
	// 时间 日志级别 哪个文件 哪一行 哪个函数
	//f.logFile.WriteString(msg)
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

	// 往文件里面写之前要做的一个检查,这个是让大小切割的
	/*
		if f.checkSplit(f.logFile) {
			f.logFile = f.splitLogFile(f.logFile)
		}

		fmt.Fprintf(f.logFile, format, args...) // 先用args替换 format里面的占位符,然后写到f.logFile里面
		fmt.Fprintln(f.logFile)                 // 加一个换行

		// 如果是error或者fatal级别日志还要记录到f.errorFile里面
		if level >= ErrorLevel {
			if f.checkSplit(f.errFile) {
				f.errFile = f.splitLogFile(f.errFile)
			}

			fmt.Fprintln(f.errFile, format)
			fmt.Fprintln(f.errFile) // 加一个换行
		}

	*/
	if f.checkSplitByHour(f.logFile) {
		// 为true,切割
		f.logFile = f.splitLogFileByHour(f.logFile)
	}

	fmt.Fprintf(f.logFile, format, args...) // 先用args替换 format里面的占位符,然后写到f.logFile里面
	fmt.Fprintln(f.logFile)                 // 加一个换行

	// 如果是error或者fatal级别日志还要记录到f.errorFile里面
	if level >= ErrorLevel {
		if f.checkSplitByHour(f.errFile) {
			// 为true,切割
			f.logFile = f.splitLogFileByHour(f.errFile)
		}
		fmt.Fprintln(f.errFile, format)
		fmt.Fprintln(f.errFile) // 加一个换行


		/* 通过日志大小来切
		if f.checkSplit(f.errFile) {
			f.errFile = f.splitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, format)
		fmt.Fprintln(f.errFile) // 加一个换行
		 */
	}

}

// 记录日志
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// close 关闭文件句柄
func (f *FileLogger) Close() {
	f.logFile.Close()
	f.errFile.Close()

}
