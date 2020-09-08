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

func transactionSQL()  {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Println("begin transaction failure, err: %v\n", err)
		return
	}

	// 执行多个SQL操作
	sqlStr1 := `update user set age=age+2 where id = ?`
	ret1, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println(err)
		return
	}

	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	sqlStr2 := `update user set age=age - 2    where id = ?`
	ret2, err := tx.Exec(sqlStr2, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Println(err)
		return
	}

	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	fmt.Println("affRow1: ", affRow1, "\t", "affRow2: ", affRow2)


	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}


func main() {
	// mysql事务

	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connection db ok ...")

	transactionSQL()
}
