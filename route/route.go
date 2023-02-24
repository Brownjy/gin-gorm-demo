package route

import (
	"github.com/gin-demo/go-project/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 测试
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	{
		// 增
		v1.POST("/todo", controller.CreateTodo)
		// 删
		v1.DELETE("/todo/:id", controller.DeleteATodo)
		// 查所有
		v1.GET("/todo/", controller.GetTodoList)
		// 改
		v1.PUT("/todo/:id", controller.UpdateATodo)
	}
	return r
}
