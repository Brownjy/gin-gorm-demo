package main

import (
	"fmt"
	"github.com/gin-demo/go-project/conf"
	"github.com/gin-demo/go-project/route"
)

func main() {
	// 初始化Config(MySql)
	conf.ConfigInit()
	// 加载路由
	r := route.Router()
	// 启动gin服务
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("gin服务启动失败...,err: ", err)
	}
}
