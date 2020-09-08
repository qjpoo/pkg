package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func main() {

	// 指定配置文件
	viper.SetConfigFile("./viper/config.yaml")
	// 指定查找配置文件的路径
	viper.AddConfigPath("./viper")
	// 读取配置信息失败
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// 监控配置文件变化
	viper.WatchConfig()

	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})
	if err := r.Run(fmt.Sprintf(":%d", viper.GetInt("port")));err != nil {
		panic(err)
	}

}
