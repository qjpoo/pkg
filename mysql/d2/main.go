package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

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

func queryDbOnce(id int) *User {
	// 使用QueryRow一定要调用 Scan方法.
	// QueryRow查询一条记录
	var u1 User
	// 1. 写查询单条记录的SQL语句
	sqlStr := `select id, name, age from user where id = ?`
	// 2. 执行
	rowObj := db.QueryRow(sqlStr, id) // 从连接池里拿一个连接出来去数据库里机查询单条记录
	// 3. 拿到结果id, name, age 给结构体对应的字段
	err := rowObj.Scan(&u1.id, &u1.name, &u1.age) // 必须要调用scan方法,因为scan里面有一个close(释放,用完之后会放到连接池里面),会释放这个连接.如果池了里面的连接不够了,新来的查询会一直等别的链接释放
	//fmt.Println(err, u1)
	fmt.Println(err, u1)
	//u1.name = "chiling48"
	return &u1
}

// 多条结果集
func queryManyDb(id int) {
	// 1. sql语句
	sqlStr := `select id, name, age from user where id >?`
	// 2. 执行
	rowObj, err := db.Query(sqlStr, id)
	if err != nil {
		log.Fatal(err)
	}
	// 3. 一定要关闭rowObj
	defer rowObj.Close()

	// 循环取值
	for rowObj.Next() {
		var u User
		err := rowObj.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}

}

/*
use test;

CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

select * from test.user;


*/

type User struct {
	id   int
	name string
	age  int
}

func main() {

	//
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connection db ok ...")

	// query one resut
	//ret := queryDbOnce(1)
	//fmt.Println(*ret)

	// query many result
	queryManyDb(1)
}
