package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {

	db, err := gorm.Open("mysql", "root:qujian123@tcp(47.98.179.41:13360)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//fmt.Println("db: ", db)
	defer db.Close()

	db.AutoMigrate(&User{})  // 创建表

	/*
	name := "chiling"
	u := User{Name: &name, Age: 48}
	ok := db.NewRecord(&u) // 判断主键是否为空,如果返回为false,说明表里面有这条数据了
	fmt.Println(ok) // true 说明是一条新的记录
	db.Create(&u)
	fmt.Println(db.NewRecord(&u)) // false
	u1 := User{Age:28} // 小王子, 28    INSERT INTO `users` (`age`) VALUES (28)
	db.Debug().Create(&u1)

	u2 := User{Name: new(string), Age:58} // 插入的是 "", 58  INSERT INTO `users` (`name`,`age`) VALUES ('',58)
	db.Debug().Create(&u2)

	u3 := User{}
	db.Debug().Create(&u3) // 小王子, ''
	 */

	name := "hello_server"
	u4 := User{Name: &name, Age: new(int64)} // hello_server, 0
	db.Debug().Create(&u4)


}

// 定义模型
type User struct {
	ID int64
	Name *string `gorm:"default:'小王子'"`
	Age *int64 `gorm:"default:0"`
}

