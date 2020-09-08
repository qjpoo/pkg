package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

var itemChan chan *item
var resultChan chan *result

// 生产者
func productor(ch chan *item) {
	// 1. 生成随机数
	var id int64
	for {
		id++
		number := rand.Int63()
		tmp := &item{
			id:  id,
			num: number,
		}
		// 2. 把随机数发送到通道中
		ch <- tmp
	}

}

// 计算一个数字每个位之合
func calc(num int64) (int int64) {
	/*
		123 % 10 = 12...3
		12 % 10 = 1 ... 2
		1 % 10 = 0 ...1
	*/
	var sum int64
	for num > 0 {
		sum = sum + num%10 // 0 + 3
		num = num / 10     // 12
	}
	return sum
}

// 消费者
func consumer(ch chan *item, resultChan chan *result) {
	for tmp := range ch { // 从ch通道里面消费
		sum := calc(tmp.num)
		// 构造一个result结构体
		retObj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retObj
	}
}

// 打印结果
func printResult(resultChan chan *result) {
	for ret := range resultChan {
		//fmt.Printf("%T\n", ret.item.id)
		fmt.Printf("id: %v, num: %v, sum: %v\n", ret.item.id, ret.item.num, ret.sum)
	}

}

func startWorker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

var wg sync.WaitGroup

func main() {
	// 生产者消费模型
	// 生产者: 产生随机数
	// 消费者: 计算每个随机数,每个位,数字的和
	//

	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 100)
	go productor(itemChan)

	startWorker(10, itemChan, resultChan)

	// 打印结果
	printResult(resultChan)

	//go consumer(itemChan, resultChan)

	/*
		rand.Seed(time.Now().UnixNano())
		//num1 := rand.Int63()
		num := rand.Intn(11) // [0-10]
		//fmt.Println(num, num1)
		var ch chan int
		ch = make(chan int, 100)
		go productor(num, ch)
	*/
}
