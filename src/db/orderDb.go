package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OrderDb() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/novice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	//关闭自动把表单名变复数
	db.SingularTable(true)
	updateOrder(db)
	queryOrder(db)
	likeQueryOrder(db)
	work(db)
	defer db.Close()
}

//事务sql
func work(db *gorm.DB) {
	tx := db.Begin()
	// 注意，一旦你在一个事务中，使用tx作为数据库句柄
	if sortQueryOrder(tx) {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
