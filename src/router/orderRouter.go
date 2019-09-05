package router

import (
	"github.com/gin-gonic/gin"
	"handle"
)

func OrderRouter() {
	r := gin.Default()
	//增加订单http://localhost:8080/creatOrder
	r.GET("/creatOrder", handle.CreateServiceHandle())
	//修改订单amount、status、file_url
	r.GET("/updateOrder", handle.UpdateServiceHandle())
	//查询订单详情
	r.GET("/queryOrder", handle.QueryServiceHandle())
	//查询订单列表，模糊查询及金额排序
	r.GET("/likeQueryOrder", handle.LikeQueryServiceHandle())

	r.Run() //:8080
}
