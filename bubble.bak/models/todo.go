package models

import (
	"pkg/bubble/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// todo  Model 增删改查

// createTodo 创建todo
func CreateTodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

// 得到列表
func GetTodo() (todoList []*Todo, err error) {
	// 查询todo表里面所有的数据
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}


// GetTodoById
func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo) // 分配一个地址
	if err = dao.DB.Where("id=?", id).First(todo).Error;err !=nil {  // First()这里面要传一个地址
		return nil, err
	}
	return
}

// UpdateToo
func UpdateTodo(todo *Todo) (err error) {
	if err = dao.DB.Save(&todo).Error; err != nil {
		return
	}
	return
}

// DeleteTodo
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
