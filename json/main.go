/*

go 语言支持 JSON 的编码 encoding 和 解码 decoding
包括内置和自定义数据类型

*/

package main

import "encoding/json"
import "fmt"
import "os"

//#下面将使用这两个结构体来演示自定义类型的编码和解码。
// 结构体 Response1，包括两个属性
type Response1 struct {
	Page   int
	Fruits []string
}

// 结构体 Response2，包括两个属性
type Response2 struct {
	// 可以在struct字段声明上使用标记来自定义编码的JSON密钥名称。
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	//首先，一些基本数据类型编码为JSON字符串的示例如下。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//这里有一些切片slice和字典map示例如下。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON包可以自动编码您的自定义数据类型。
	// 它只包含编码输出中的导出字段，默认情况下将这些名称用作JSON键。
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 检查上面的“Response2”的定义，看看这样的标签的例子。
	// 字段的大小写不一样呀
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 现在让我们看看将JSON数据解码为Go值。
	// 这是一个通用数据结构的示例。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 建立一个用于JSON包放置解码数据的变量dat。
	// 这个`map [string]interface{}`将一个字符串映射到任意数据类型。
	var dat map[string]interface{}

	// 这是实际的解码，并检查相关的错误是否存在，并打印
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码映射中的值，我们需要将它们转换为适当的类型。
	// 例如，我们将`num`中的值转换为预期的`float64`类型。
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套数据需要一系列强制转换。
	strs := dat["strs"].([]interface{}) //转换为任意数据类型
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们还可以将JSON解码为自定义数据类型。
	// 这样做的好处是可以为我们的程序增加额外的类型安全性
	// 并且在访问解码数据时不需要类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 在上面的例子中，
	// 我们总是使用字节和字符串作为标准输出中数据和JSON表示之间的中介。
	// 我们也可以直接将JSON编码流式传输到`os.Writer，就像`os.Stdout`甚至是HTTP响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
