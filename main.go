package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = dreams
	//p.dreams[0] ="不吃饭"
}

func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?

	str := "2006.01.02 15:04:05"
	sliceStr := strings.Split(str, " ")
	fmt.Println(sliceStr[0])
	fmt.Printf("%T, %v\n", sliceStr[1], sliceStr[1])
	h := strings.Split(sliceStr[1], ":")
	fmt.Printf(h[0])

	fmt.Println("\n------------------")
	//t := time.Now()
	//fmt.Println(t.Truncate(time.Hour * 24))
	t := time.Now()
	fmt.Println(t.Truncate(time.Minute * 24))

	ko := `filePath="/root/123.log"`
	koNew := strings.ReplaceAll(ko, "\"", "")
	fmt.Println(koNew)

	// 多个分隔符切割
	pathStr := `filePath="/root/123.log"`
	fieldsFunc := strings.FieldsFunc(pathStr, func(r rune) bool {
		return r == '"' || r == '=' || r == '/'
	})
	fmt.Println(fieldsFunc)
	words := strings.FieldsFunc(pathStr, func(r rune) bool { return strings.ContainsRune("=\"/", r) })
	fmt.Println(words)

	fmt.Println("-----------------")
	f1(hoo,ooh)
	fmt.Println(ooh)


/*	fmt.Println("-----------------")
	var po *Person
	po.name= "xxx"
	fmt.Println(po)
	po = new(Person)
	fmt.Println(po)
*/	/*
	po =&Person{
		name: "qujianjian",
		age: 18,
	}
	fmt.Println(po)  // &{qujianjian 18 []}
	fmt.Println(*po)   // {qujianjian 18 []}
	 */

	fmt.Println("--------变量的作用域------")
	fmt.Println("before varn: ", varn)
	setVarn(100)
	fmt.Println("after varn: ", varn)
	varn = 122
	fmt.Println("change varn: ", varn)
	getS()


	// mashal
	m := Message{"Alice", "Hello", 1294706395881547000}
	b1, _ := json.Marshal(m)
	fmt.Println(string(b1))
	//fmt.Printf("type: %T, value: %#v", b1, b1)

	fmt.Println("------->")
	b0 := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	//var m1 map[string]interface{}
	m1 := make(map[string]interface{})
	json.Unmarshal(b0, &m1)

	if v8, ok := m1["Parents"].([]interface{});!ok{
		fmt.Println("err")
	}else {
		fmt.Println(v8)
		for _, value0 := range v8 {
			k9 := value0.(string)
			fmt.Printf("%T, %v\n", k9, k9)
		}
	}
	for _, v := range m1 {
		//fmt.Println("------>: ", v)
		switch  vv := v.(type){  // //vv表示v1接口转换成的值
		case float64:
			fmt.Println("===========float64: ", vv)
		case string:
			fmt.Println("string: ", vv)
		case []interface{}:
			//fmt.Println(k, "is []interface ...")
			for i, u := range vv {
				fmt.Println(i, u)
		}

		}
	}


}

/*
正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。

func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
 */


var hoo = 10
var ooh int

func f1(hoo,ooh int)  {
	ooh = 101
	fmt.Println(hoo, ooh)
}

var varn = 199

func setVarn(i int)  {
	varn = i
	fmt.Println("setVarn func varn: ", varn)
}

func getS()  {
	fmt.Println("getS func varn: ", varn)
}


type Message struct {
	Name string
	Body string
	Time int64
}



