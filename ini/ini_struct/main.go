package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"reflect"
	"time"
)

type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

func main() {
	cfg, err := ini.Load("./ini/ini_struct/info.ini")
	if err != nil {
		fmt.Println(err)
	}
	//p := new(Person)
	p := Person{}
	if err = cfg.MapTo(&p); err != nil {
		fmt.Printf("%v", fmt.Sprintf("err: %v", err))
	}
	//fmt.Println(p)

	t := reflect.TypeOf(p)
	fmt.Println(t) // *main.Person
	v := reflect.ValueOf(p)
	fmt.Println(v) // &{Unknwon 21 true 1993-01-01 20:17:05 +0000 UTC {Hi is a good man! [HangZhou Boston]}}
	for k := 0; k < t.NumField(); k++ {
		fmt.Println(t.Field(k).Name, v.Field(k).Interface())
	}
}
