package service

import (
	"model"
	"testing"
)

//测试填加订单
func TestCreateService(t *testing.T) {
	demoOrder := model.DemoOrder{OrderId: "1", UserName: "523", Amount: 43, Status: "1", FileURL: "22"}
	CreateService(&demoOrder)
}

//测试修改订单amount、status、file_url
func TestUpdateService(t *testing.T) {
	demoOrder := model.DemoOrder{OrderId: "1", UserName: "24", Amount: 131, Status: "3", FileURL: "4"}
	UpdateService(&demoOrder)
}

//测试查询订单详情
func TestQueryService(t *testing.T) {
	demoOrder := model.DemoOrder{OrderId: "1"}
	QueryService(&demoOrder)
}

//测试查询订单列表，模糊查询及金额排序
func TestLikeQueryService(t *testing.T) {
	var demoOrder []model.DemoOrder
	LikeQueryService(&demoOrder)
}

//测试demo_order 所有数据以excel形式导出来(可以下载)
func TestExcelExportService(t *testing.T) {
	ExcelExportService()
}
