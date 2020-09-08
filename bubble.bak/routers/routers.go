package routers

import (
	"github.com/gin-gonic/gin"
	"pkg/bubble/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// 静态文件路径
	r.Static("/static", "./bubble/static")

	// 指定html文件位置
	r.LoadHTMLGlob("./bubble/templates/*")

	r.GET("/", controller.IndexHander)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodo)


		// 修改
		v1Group.PUT("/todo/:id", controller.ModifyTodo)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}


	return r

}
