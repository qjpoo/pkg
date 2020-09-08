package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	// redis

	conn, err := redis.Dial("tcp", "47.98.179.41:16379", redis.DialPassword("quja@123"))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	/*
	info, err := conn.Do("info")
	if err != nil {
		fmt.Println(err.Error())
	}
	//a := info
	//fmt.Printf("%T, %T\n", info, a)
	//infoStr := fmt.Sprintf("%v", info)
	//fmt.Println(infoStr)
	//fmt.Println()
	v, ok := info.([]uint8)
	if ok {
		//fmt.Println("info: ", v)
		fmt.Println(string(v))
	}else {
		fmt.Println("error ...")
	}
	*/

	// string
	_, err = conn.Do("set", "name", "qujian")
	if err != nil {
		fmt.Println(err.Error())
	}
	str, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)


	// int
	_, err = conn.Do("set", "age", 20)
	if err != nil {
		fmt.Println(err.Error())
	}
	age, err := redis.Int(conn.Do("get", "age"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(age)


	// hash
	_, err = conn.Do("hset", "books","abc", 200)
	if err != nil {
		fmt.Println(err.Error())
	}
	num, err := redis.Int(conn.Do("hget", "books", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	// 批量set
	_, err = conn.Do("mset", "apple", 100, "banana", 20)
	if err != nil {
		fmt.Println(err.Error())
	}
	n, err := redis.Ints(conn.Do("mget", "apple", "banana"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	// 过期时间
	_, err = conn.Do("expire", "age", 5)
	if err != nil {
		fmt.Println(err)
		return
	}

}
