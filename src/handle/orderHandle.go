package handle

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"model"
	"service"
)

//填加订单
func CreateServiceHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		demoOrder := model.DemoOrder{OrderId: CreateRandomNumber(8), UserName: "523", Amount: 43, Status: "1", FileURL: "22"}
		flag := service.CreateService(&demoOrder)
		var success string
		if flag {
			success = "操作成功"
		} else {
			success = "操作失败"
			demoOrder = model.DemoOrder{}
		}
		c.JSON(200, gin.H{
			"success": success,
			"message": demoOrder,
		})
	}
}

//修改订单amount、status、file_url
func UpdateServiceHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		demoOrder := model.DemoOrder{OrderId: "11", UserName: "24", Amount: 131, Status: "3", FileURL: "4"}
		flag := service.UpdateService(&demoOrder)
		var success string
		if flag {
			success = "操作成功"
		} else {
			success = "操作失败"
			demoOrder = model.DemoOrder{}
		}
		c.JSON(200, gin.H{
			"success": success,
			"message": demoOrder,
		})
	}
}

//查询订单详情
func QueryServiceHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		demoOrder := model.DemoOrder{OrderId: "77"}
		flag := service.QueryService(&demoOrder)
		var success string
		if flag {
			success = "操作成功"
		} else {
			success = "操作失败"
			demoOrder = model.DemoOrder{}
		}
		c.JSON(200, gin.H{
			"success": success,
			"message": demoOrder,
		})
	}
}

//查询订单列表，模糊查询及金额排序
func LikeQueryServiceHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var demoOrder []model.DemoOrder
		flag := service.LikeQueryService(&demoOrder)
		var success string
		if flag {
			success = "操作成功"
		} else {
			success = "操作失败"
		}
		c.JSON(200, gin.H{
			"success": success,
			"message": demoOrder,
		})
	}
}

//生成指定位数随机数
func CreateRandomNumber(len int) string {
	var numbers = []byte{0, 1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= len; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {

		}
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}
