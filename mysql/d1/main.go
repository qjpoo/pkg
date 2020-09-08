package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// mysql

	// 连接用户名 data source name
	dsn := "root:qujian123@tcp(47.98.179.41:13360)/bubble"
	db, err := sql.Open("mysql", dsn)  // 不会提供校验数据是否正确

	if err != nil {
		fmt.Println(err)
	}

	e := db.Ping()
	if err !=nil {
		fmt.Println(e)
	}else {
		fmt.Println("connection db sucessfull ...")
	}
}
