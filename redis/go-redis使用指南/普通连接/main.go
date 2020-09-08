package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	ExampleNewClient()
	ExampleClient()
}

//ping pong测试
func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "115.159.128.112:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

//set和get测试
func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "115.159.128.112:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//第三个参数是过期时间
	err := client.Set("name", "taoshihan", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", val)
	//检测key是否存在
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
