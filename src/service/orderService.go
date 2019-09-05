package service

import (
	"db"
	"fmt"
	"model"
)

//填加订单
func CreateService(demoOrder *model.DemoOrder) bool {
	db := db.GetDb()
	//事务
	tx := db.Begin()
	err := db.Create(&demoOrder).Error
	// 注意，一旦你在一个事务中，使用tx作为数据库句柄
	if err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	defer db.Close()
	return isSuccess(err)
}

//修改订单amount、status、file_url
func UpdateService(demoOrder *model.DemoOrder) bool {
	db := db.GetDb()
	//执行成功返回nil，失败返回失败信息
	err := db.Model(&demoOrder).Where("order_id = ?", demoOrder.OrderId).Update(demoOrder).Error
	defer db.Close()
	return isSuccess(err)
}

//查询订单详情
func QueryService(demoOrder *model.DemoOrder) bool {
	db := db.GetDb()
	err := db.Where("order_id = ?", demoOrder.OrderId).Find(&demoOrder).Error
	defer db.Close()
	return isSuccess(err)
}

//查询订单列表，模糊查询及金额排序
func LikeQueryService(demoOrder *[]model.DemoOrder) bool {
	db := db.GetDb()
	err := db.Where("user_name like ?", "%2%").Order("amount").Find(&demoOrder).Error
	defer db.Close()
	return isSuccess(err)
}

//成功执行则返回true
func isSuccess(err error) bool {
	var flag bool
	if err != nil {
		flag = false
		fmt.Println("操作失败", err)
	} else {
		flag = true
	}
	return flag
}
