package main
import(
	"fmt"
)
func main(){
	a := 1
	fmt.Println("a int: ",&a)
	aStr := "a"
	fmt.Println("aStr: ", &aStr)
	aSlice := []int{1, 2,3}
	fmt.Println("aSlice: ", &aSlice)   // aSlice:  &[1 2 3]
	arr:=[...]string{"aa","bb","cc","vv"}
	fmt.Println("outside: ",&arr)  // outside:  &[aa bb cc vv]
	test1(&arr)
	fmt.Println(&arr)  // &[hahaha bb cc vv]
	//这样取到的是变量的地址，但是保存变量的地址需要的是一个指针类型

	//var arrP *[4]string=&arr

	//fmt.Printf("%v\n",&arrP)

	// fmt.Println(*(&arr))
	// 对指针变量取他的值
	//test1(arrP)
}


func test1(arr *[4]string){
	//(*arr)[0]="hahaha"
	arr[0]="haha"
	fmt.Println("*arr: ",*arr)  //  [haha bb cc vv]
	fmt.Println("arr: ", arr)  // &[haha bb cc vv]
	fmt.Println("&arr", &arr)  // 0xc000148018
	// 传入的是指针，先用取值符找到数组
	// arr后面才是真正传入的类型 我要传入的是一个有4个元素的数组，而且数组的值是字符串

}