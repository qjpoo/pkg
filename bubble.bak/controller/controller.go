package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/bubble/models"
)

/*
url 		==> 		controller 		==> 		login 		==> model
请求来了		==>			控制器		    ==>  		业务逻辑     ==> 模型层的增删改查
*/
func IndexHander(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项,点击提交
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)

	// 2. 存入数据库
	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 20001,
			"msg":  "ok",
			"data": todo,
		})
	}
}

func GetTodo(c *gin.Context) {
	// 查询todo表里面所有的数据
	todoList, err := models.GetTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 20002,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}

}

func ModifyTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}

	todo, err := models.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 拿到数据
	c.BindJSON(&todo)

	if err = models.UpdateTodo(todo);err !=nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}

	//var todo Todo
	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": id + "已删除"})
	}
}
