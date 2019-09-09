package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"novice/model"
)

func GetDb() *gorm.DB {
	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/novice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	//关闭自动把表单名变复数
	db.SingularTable(true)
	var demoOrder model.DemoOrder
	//检查是否存在该表，没有则创建
	if !db.HasTable("demo_order") {
		// 为模型`DemoOrder`创建表
		db.CreateTable(&demoOrder)
	}
	return db
}
