package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 连接mysql
	db, err := gorm.Open("mysql", "root:qujian123@tcp(47.98.179.41:13360)/bubble")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	fmt.Println("ok ...")
}
