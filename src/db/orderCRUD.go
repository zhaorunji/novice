package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"model"
)

func orderCRUD() {

}

//修改更新订单数据库
func updateOrder(db *gorm.DB) {
	order := model.Order{Amount: 121, Status: "3", FileURL: "4"}
	db.Table("demo_order").Model(&order).Where("order_id = ?", 1).Update(order)
}

//查询订单列表
func queryOrder(db *gorm.DB) {
	var array = []model.Order{}
	db.Table("demo_order").Find(&array)
	fmt.Println("查询订单列表:", array)
}

//模糊查询订单
func likeQueryOrder(db *gorm.DB) {
	var array []model.Order
	db.Table("demo_order").Where("user_name like ?", "%2").Find(&array)
	fmt.Println("模糊查询订单:", array)
}

//订单金额排序
func sortQueryOrder(db *gorm.DB) bool {
	var array []model.Order
	err := db.Table("demo_order").Order("amount").Find(&array)
	fmt.Println("订单金额排序:", array)
	flag := err != nil
	return flag
}
