package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func main() {
	// 只会修改默认表的前缀
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return "prefix_" + defaultTableName;
	//}

	db, err := gorm.Open("mysql", "root:qujian123@tcp(47.98.179.41:13360)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//fmt.Println("db: ", db)
	defer db.Close()

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//db.SingularTable(true)

	// 使用User结构体创建名为`deleted_users`的表
	//db.Table("deleted_users").CreateTable(&User{})

	db.AutoMigrate(&User{}) // 自动创建表
	db.AutoMigrate(&Animal{}) // 自动创建表




}

// 定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"`// 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}


// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

/*
// Animal结构体,会创建xxoo表
func (Animal) TableName() string {
	return "xxoo"
}
 */
