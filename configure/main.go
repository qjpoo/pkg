package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Config struct {
	FilePath string `conf:"file_path"`
	FileName string `conf:"file_name"`
	MaxSize  int64  `conf:"max_size"`
}

// 从conf文件中读取内容赋值给结构体指针
func parseConf(fileName string, result interface{}) (err error) {
	// 前提条件, result必须是一个ptr
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	if t.Kind() != reflect.Ptr {
		err = errors.New("result必须是一个指针 ")
		return
	}

	// result传进来的值的类型,不是结构体也不行.就是要指针里面存的值(内存地址)指向的要是一个结构体
	fmt.Println("t.elem().kind: ", t.Elem().Kind(), "t.kind: ", t.Kind(), "v.elem.kind: ", v.Elem().Kind()) // t.elem().kind: struct  t.kind:  ptr  v.elem.kind:  struct
	if v.Elem().Kind() != reflect.Struct {
		err = errors.New("result必须是一个结构体指针")
		return
	}

	// 1. 打开文件
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		//fmt.Println(err)
		//errors.New("打开配置文件失败")
		err = fmt.Errorf("打开配置文件%s失败", fileName)
		return err
	}
	// 2. 将读取的文件数据按照行来分割
	// 得到一个行的切片
	lineSlice := strings.Split(string(data), "\r\n")
	fmt.Println(lineSlice)

	// 一行一行的解析
	for index, line := range lineSlice {
		line = strings.TrimSpace(line) // 去掉字符串首尾的空白
		line = strings.ReplaceAll(line, "\"", "")
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// 是注释或者是空行
			continue
		}
		// 解析看是不是正确的配置项
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 {
			err = fmt.Errorf("第%d行语法错误", index+1)
			return
		}
		// 按照=切割一行,左边是Key, 右边是value
		key := line[:equalIndex]
		value := line[equalIndex+1:]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if len(key) == 0 {
			err = fmt.Errorf("key为空 ...")
			return
		}
		//fmt.Println(value)

		// 利用反射, 给result赋值
		// 遍历结构体的每一个字段和key比较, 匹配上就把value赋值
		for i := 0; i < t.Elem().NumField(); i++ {
			field := t.Elem().Field(i)
			tag := field.Tag.Get("conf")
			if key == tag {
				// 匹配上就赋值
				// 拿到每个字段的类型
				fieldType := field.Type
				//fmt.Println(fieldType)
				switch fieldType.Kind() {
				case reflect.String:
					// 根据字段名,拿到字段的值
					fieldValue := v.Elem().FieldByName(field.Name)
					// 将配置文件中读取的值value设置给结构体
					fieldValue.SetString(value)
					v.Elem().Field(i).SetString(value)
				case reflect.Int64, reflect.Int32, reflect.Int:
					value64, _ := strconv.ParseInt(value, 10, 64) // 把字符串转化成 10进制的Int64
					v.Elem().Field(i).SetInt(value64)
				}

			}

		}
	}
	return

}

func main() {
	// 解析配置文件

	// 2. 读取内容

	// 3. 一行一行读取内容, 根据TAG找结构体里面的对应字段

	// 4. 找到要赋值

	var c = &Config{} // 用来存储读取的数据
	err := parseConf("./configure/system.conf", c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", c) // 解析之后的数据

	//	readFile("configure/system.conf")
	/*
		ini配置文件的解析
		[mysqld]
		host = 127.0.0.1
		username = xxoo
		password = xxoo
		port = 3306

		[redis]
		host = 127.0.0.1
		password = xxoo
		port = 6379

		;或者#都是注释

	*/

}

/*
func readFile(filePath string) {
	//var file *os.File
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		//l := make([]byte,1024)
		l, _, err := rd.ReadLine()
		//fmt.Println(string(l))
		t := strings.FieldsFunc(string(l), func(r rune) bool {
			return r == '=' || r == '"' || r == ' '
		})
		for k, v := range t {
			fmt.Printf("key: %d value: %v\n", k, v)
		}
		if err != nil || err == io.EOF {
			break
		}
	}

}
*/
