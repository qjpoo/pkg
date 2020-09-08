package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pkg/bubble/dao"
	"pkg/bubble/models"
	"pkg/bubble/routers"
)

func main() {
	// 创建数据库
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close() // 关闭数据库

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	routers.SetupRouter().Run(":8000")
}
