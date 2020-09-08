package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pkg/bubble/setting"
)

var DB *gorm.DB // 是一个连接池对象

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	//fmt.Println("cfg: ", cfg)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	//dsn := "root:qujian123@tcp(47.98.179.41:13360)/bubble"
	fmt.Println(dsn)
	DB, err = gorm.Open("mysql", dsn)
	fmt.Println("DB: ", DB, "err:", err)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
