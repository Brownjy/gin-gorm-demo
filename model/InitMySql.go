package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MySQLInit(dsn string) {
	//1.连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败,err: ", err)
		return
	}
	fmt.Println("连接数据库成功！")
	DB = db
	//2.迁移表
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		fmt.Println("迁移数据库失败！")
		return
	}
}
