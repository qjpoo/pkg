package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := [3]int{1, 2, 3}
	x1 := &x
	fmt.Println(x1)
	func(arr *[3]int) {
		fmt.Println(arr)
		(*arr)[0] = 7
		fmt.Println(arr) // &[7 2 3]
	}(&x)
	fmt.Println(x) // [7 2 3]

	var data interface{} = "great"

	if res, ok := data.(int); ok {
		fmt.Println("[is an int], data: ", res)
	} else {
		fmt.Println("[not an int], data: ", data) // [not an int], data:  great
	}

	/*
	m3 := map[string]data1{
		"x": {"Tom"},
	}
	m3["x"].name = "Jerry"
	 */

	y1 := []data1{{name: "xxoo"},{name: "xxaa"}}
	y1[0].name="xx11"
	fmt.Println(y1)

<<<<<<< HEAD

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
=======
>>>>>>> 72fb6378bb59539679f1f9cb6d02b6fb6e0afa61

}

type data1 struct {
	name string
}
<<<<<<< HEAD


type Message struct {
	Name string
	Body string
	Time int64
}



=======
>>>>>>> 72fb6378bb59539679f1f9cb6d02b6fb6e0afa61
