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

func prepareDB() {
	sqlStr := `select id, name, age from user where id > ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
	}

}

// 预插入
func prepareInsert()  {
	sqlStr := `insert into user(name, age) values(?, ?)`
	stmt, err := db.Prepare(sqlStr) // 把SQL语句先发给服务器预处理
	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	var m = map[string]int{
		"王强": 20,
		"少帅": 32,
		"周勇": 34,
	}
	for k, v := range m {
		_, err = stmt.Exec(k, v)  // 把值发给服务器
		if err != nil {
			fmt.Println(err )
		}
		fmt.Println("insert into user success ....")
	}

}

// 预更新
func prepareUpdate()  {
	sqlStr := `update user set age= ? where id = ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(43, 5)
	if err != nil {
		fmt.Println(err )
	}

	fmt.Println("update  user success ....")
}


// 预删除
func prepareDelete()  {
	sqlStr := `delete from user where id = ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(6)
	if err != nil {
		fmt.Println(err )
	}

	fmt.Println("delete  user success ....")
}
func main() {
	// mysql预处理

	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connection db ok ...")

	prepareDB()

	prepareInsert()

	prepareUpdate()

	prepareDelete()

	fmt.Println("----------------------")
	prepareDB()

}
