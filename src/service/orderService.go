package service

import (
	"db"
	"fmt"
	"github.com/tealeg/xlsx"
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

//将demo_order 所有数据以excel形式导出来(可以下载)
func ExcelExportService() bool {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("demo_order_list")
	var demoOrders []model.DemoOrder
	//获取demo_order数据
	LikeQueryService(&demoOrders)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		//创建列名
		row := sheet.AddRow()
		orderId := row.AddCell()
		orderId.SetValue("order_id")
		userName := row.AddCell()
		userName.SetValue("user_name")
		amount := row.AddCell()
		amount.SetValue("amount")
		status := row.AddCell()
		status.SetValue("status")
		fileUrl := row.AddCell()
		fileUrl.SetValue("file_url")
		//添加行内容
		for _, demoOrder := range demoOrders {
			row := sheet.AddRow()
			orderId := row.AddCell()
			orderId.SetValue(demoOrder.OrderId)
			userName := row.AddCell()
			userName.SetValue(demoOrder.UserName)
			amount := row.AddCell()
			amount.SetValue(demoOrder.Amount)
			status := row.AddCell()
			status.SetValue(demoOrder.Status)
			fileUrl := row.AddCell()
			fileUrl.SetValue(demoOrder.FileURL)
		}
	}
	//写入的文件
	err = file.Save("./file/excelDemo.xlsx")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
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
