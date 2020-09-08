package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int
	name string
	age  int
}

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	dsn := "root:qujian123@tcp(47.98.179.41:13360)/test"
	db, err = sql.Open("mysql", dsn) // 这里不能写:= ,否则db就是一个局部变量了,最上面申明的全局变量没有使用
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置数据库连接池最大数
	db.SetMaxOpenConns(10)

	// 最大空闲连接数
	db.SetMaxIdleConns(3)
	return nil
}

// 插入数据
func insertDB()  {
	sqlStr := `insert into user(name, age) value (?, ?)`
	ret, err := db.Exec(sqlStr, "王五", 56)
	if err != nil {
		fmt.Println(err)
	}

	// 新插入数据的ID
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("insert success, the id is %d\n", id)
}

func updateDB()  {
	sqlStr := `update user set age=? where id = ?`
	ret, err := db.Exec(sqlStr,20, 2)
	if err != nil {
		fmt.Println(err)
	}

	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("update success, affected rows: %d\n", n)

}

func deleteDB()  {
	sqlStr := `delete from user where id = ?`
	ret, err := db.Exec(sqlStr, 8)
	if err != nil {
		fmt.Println(err)
	}
	n, err := ret.RowsAffected()
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Printf("delete ok, affected rows: %d\n", n)
}
func main() {
	// SQL的插入

	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connection db ok ...")

	// 插入数据
	insertDB()

	// 更新数据
	updateDB()

	// 删除数据
	deleteDB()

}
