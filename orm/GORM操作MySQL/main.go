package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// * DB
	db, err := gorm.Open("mysql", "root:qujian123@tcp(47.98.179.41:13360)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//fmt.Println("db: ", db)
	defer db.Close()

	// 创建表,自动迁移,(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//u1 := UserInfo{1, "chiling", "男", "golang"}
	//db.Create(&u1)
	//u2 := UserInfo{2, "qujian", "男", "movie"}
	//db.Create(u2)

	// 查询
	var u UserInfo
	db.First(&u)  // 要传一个结构体地址,不然没有办法修改的
	fmt.Println(u)  // 只有一条记录   {1 chiling 男 golang}

	// update
	db.Model(&u).Update("gender", "女") // 前面是字段名 后面是字段的值
	db.Model(&u).Update("hobby", "看电影") // 前面是字段名 后面是字段的值
	fmt.Println(u)


	// delete
	db.Delete(&u)

}

// 它会把UserInfo ==> user_infos表
// 用户结构表 ==> 对应表
type UserInfo struct {
	ID     uint8
	Name   string
	Gender string
	Hobby  string
}
