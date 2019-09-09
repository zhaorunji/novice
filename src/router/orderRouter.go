package router

import (
	"github.com/gin-gonic/gin"
	"handler"
)

func OrderRouter() {
	r := gin.Default()
	//增加订单http://localhost:8080/creatOrder
	r.GET("/creatOrder", handler.CreateServiceHandler())
	//修改订单amount、status、file_url
	r.GET("/updateOrder", handler.UpdateServiceHandler())
	//查询订单详情
	r.GET("/queryOrder", handler.QueryServiceHandler())
	//查询订单列表，模糊查询及金额排序
	r.GET("/likeQueryOrder", handler.LikeQueryServiceHandler())
	//文件上传，并更新file_url
	r.POST("/fileUpLoad", handler.FileUpLoadServiceHandler())
	//文件下载
	r.GET("/fileDownLoad", handler.FileDownLoadServiceHandler())
	//将demo_order 所有数据以excel形式导出来(可以下载)
	r.GET("/excelExport", handler.ExcelExportServiceHandler())

	r.Run(":8080")
}
