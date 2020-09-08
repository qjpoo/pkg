package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Port int `config:"port"`
	Version string `config:"version"`
}

var (
	conf = new(Config)
)
func main() {
	// 配置文件和结构体对应上

	// 指定配置文件路径
	viper.SetConfigFile("./viper/config.yaml")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将读取的配置信息保存至全局变量conf
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("------------------>: ", conf)

	// 监控配置文件变化
	viper.WatchConfig()

	// 配置文件发生变化要同步到全局变量conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件有变化 ...")
		if err := viper.Unmarshal(conf); err != nil {
			fmt.Println(err)
		}
	})

	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.String(200, conf.Version)
	})

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil{
		panic(err)
	}


}
