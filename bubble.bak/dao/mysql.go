package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB // 是一个连接池对象

func InitMysql() (err error) {
	dsn := "root:qujian123@tcp(47.98.179.41:13360)/bubble"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return DB.DB().Ping()
}