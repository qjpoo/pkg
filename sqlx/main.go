package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type User struct {
	Id   int  // 一定要首字母是大写,因为db.Get()第一个变量是通过反射拿到的类型
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "root:qujian123@tcp(47.98.179.41:13360)/test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	return 
}

// 单行
// 查询单条数据示例
func queryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := db.Get(&u, sqlStr, 2)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id: %d name: %s age: %d\n", u.Id, u.Name, u.Age)
	//fmt.Printf("u: %+#v\n", u)
}

// 多行
func queryRows()  {
	sqlStr := `select id, name, age from user where id > ?`
	var u []User // 一定要是一个结构体的切片
	err := db.Select(&u, sqlStr, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range u {
		fmt.Printf("id: %d, name: %s, age: %d\n", v.Id, v.Name, v.Age)
	}
}

// get
func getResult()  {
	var count int64
	sqlStr := `select count(*) from user`
	e := db.QueryRow(sqlStr).Scan(&count)
	if e !=nil {
		fmt.Println(e)
	}
	fmt.Println(count)

}
func main() {
	// sqlx
	
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// 查询单行数据
	queryRow()


	// 多行
	queryRows()

	//
	getResult()




}
