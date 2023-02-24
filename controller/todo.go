package controller

import (
	"fmt"
	"github.com/gin-demo/go-project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建一个todo
func CreateTodo(c *gin.Context) {
	//1.从请求中把数据拿出来
	var todo model.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		fmt.Println("接收数据失败！")
		return
	}
	//2.存入数据库
	err = model.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg":  "success",
		//	"data": todo,
		//})
	}
}

func UpdateATodo(c *gin.Context) {
	// 解析URL参数
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	// 根据参数获取todo
	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 更新todo
	c.BindJSON(&todo)
	if err = model.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	// 获取所有todoList
	todoList, err := model.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	}
	err := model.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "delete_success",
			"data": "",
		})
	}
}
