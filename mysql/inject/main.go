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


// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Println(sqlStr) // select id, name, age from user where name='xxoo' or 1 = 1 #' #后面的就注释了
	var u []User // 一定要是一个结构体的切片
	err := db.Select(&u, sqlStr )
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range u {
		fmt.Printf("id: %d, name: %s, age: %d\n", v.Id, v.Name, v.Age)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connection db ok ...")

	//sqlInjectDemo("chiling")


	//sqlInjectDemo("xxoo' or 1 = 1 #")

	//sqlInjectDemo("xxx' or (select count(*) from user) <10 #")

	sqlInjectDemo("xxx' union select * from user #")
}
