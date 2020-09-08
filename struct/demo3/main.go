package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/unknwon/com"
	"io"
	"net/http"
	"time"
)

//空接口
func test1() {
	a := make(map[string]interface{}, 20)
	a["name"] = "haha"
	a["age"] = 20
	a["merried"] = true
	a["hobby"] = []string{"喝", "跳", "rap"}
	fmt.Printf("type:%T v:%#v\n", a["hobby"], a["hobby"])
	v, ok := a["hobby"].([]string)
	if !ok {
		fmt.Println("type is not map.") //type:[]string v:[]string{"喝", "跳", "rap"}
	}

	fmt.Println(v[0]) //喝

}

//接口作为函数的参数
func test2(a interface{}) {
	fmt.Printf("type:%T value:%v \n", a, a)
}

type resParamData struct {
	Code int         `json:"code"`
	Msg  string      `json:"Msg"`
	Data interface{} //方法一
}
type userinfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test3() {
	var m = make(map[string]interface{}, 0)
	m["name"] = "test"
	m["age"] = 20
	var res resParamData
	res.Code = 200
	res.Msg = "ok"
	res.Data = m

	fmt.Println(res.Data)
	value, ok := res.Data.(map[string]interface{})
	if !ok {
		fmt.Println("It's not ok for type Order")
		return
	}
	fmt.Println("The value is ", value["name"])
	fmt.Println("The value is ", value["age"])
	fmt.Printf("type:%T value:%v\n", value["age"], value["age"])

}
func test4() {
	var uinfo = userinfo{
		Name: "lisi",
		Age:  18,
	}
	var resdata resParamData
	resdata.Code = 200
	resdata.Msg = "ok"
	resdata.Data = uinfo
	fmt.Println(resdata.Data)
	value, ok := resdata.Data.(userinfo)
	if !ok {
		fmt.Println("It's not ok for type Order")
		return
	}
	fmt.Println("The value is ", value.Name)
	fmt.Println("The value is ", value.Age)
	fmt.Printf("type:%T value:%v\n", value.Name, value.Name)
	fmt.Printf("type:%T value:%v\n", value.Age, value.Age)

}

type PostresData2 struct {
	Data   interface{} //方法一
	Errmsg string      `json:"errmsg"`
	Errno  int         `json:"errno"`
}

func test5() {
	var d PostresData2
	url := "https://api.ibanana.club/select/major/list_by_key?page=1&row=10&major_name=工程"
	res := SetGet(url)  // {"data":[],"errmsg":"ok","errno":200}
	_ = json.Unmarshal([]byte(res), &d)
	fmt.Println(1111111111)
	fmt.Printf("type:%T value:%#v \n", d, d)
	fmt.Printf("type:%T value:%#v \n", d.Errno, d.Errno)
	fmt.Printf("type:%T %s value:%#v \n", d.Data,"<--->", d.Data)
	value, ok := d.Data.([]interface{})
	if !ok {
		fmt.Println("It's not ok for type major")
		return
	}
	for _, v := range value {
		m, ok := v.(map[string]interface{})
		if !ok {
			fmt.Println("It's not ok for type m")
			return
		}
		id := com.StrTo(com.ToStr(m["id"])).MustInt64()
		fmt.Printf("type:%T value:%v\n", id, id)                           //type:int64 value:20
		fmt.Printf("type:%T value:%v\n", m["major_name"], m["major_name"]) //type:string value:交通工程
		//改变期值
		if id == 19 {
			m["id"] = 100
			m["major_name"] = "edit value22"
		}
	}
	fmt.Println("GetTest")
	fmt.Println(d)
	//httpext.SuccessExt(ctx, d)
}

func main() {
	test5()
	/*
	1111111111
	type:main.PostresData2 value:main.PostresData2{Data:[]interface {}{}, Errmsg:"ok", Errno:200}
	type:int value:200
	type:[]interface {} value:[]interface {}{}
	GetTest
	{[] ok 200}
	 */

	//test4()
	/*
	{lisi 18}
	The value is  lisi
	The value is  18
	type:string value:lisi
	type:int value:18
	 */

	//test3()
	/*
	map[age:20 name:test]
	The value is  test
	The value is  20
	type:int value:20
	 */

	//test1()
	//a := 12
	//test2(a)  // type:int value:12
	//test2(nil)  // type:<nil> value:<nil>
	//test2(false)  // type:bool value:false
}

// Get ... 发送请求 ...
// url：         请求地址
// response：    请求返回的内容
func SetGet(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}