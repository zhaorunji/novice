package handler

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"novice/model"
	"novice/service"
	"path/filepath"
)

//填加订单
func CreateServiceHandler() gin.HandlerFunc {
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
func UpdateServiceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		demoOrder := model.DemoOrder{OrderId: "1", UserName: "24", Amount: 131, Status: "3", FileURL: "4"}
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
func QueryServiceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		demoOrder := model.DemoOrder{OrderId: "1"}
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
func LikeQueryServiceHandler() gin.HandlerFunc {
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

//文件上传，并更新file_url
func FileUpLoadServiceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var demoOrder model.DemoOrder
		var flag bool
		var uploadFile string
		//对应input type="file 的name
		file, err := c.FormFile("file")
		if err != nil {
			flag = false
		} else {
			uploadFile = "./file/" + file.Filename
			// 上传文件到指定的路径
			if err := c.SaveUploadedFile(file, uploadFile); err != nil {
				flag = false
			} else {
				//获取当前文件路径
				dir, _ := filepath.Abs("file")
				fileURL := dir + "\\" + file.Filename
				demoOrder = model.DemoOrder{OrderId: "1", UserName: "24", Amount: 131, Status: "3", FileURL: fileURL}
				flag = service.UpdateService(&demoOrder)
			}
		}
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

//文件下载
func FileDownLoadServiceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		downFile(c, "./file/upLoadFileTest.PNG", "upLoadFileTest.PNG")
	}
}

//将demo_order 所有数据以excel形式导出来(可以下载)
func ExcelExportServiceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		flag := service.ExcelExportService()
		if flag {
			downFile(c, "./file/excelDemo.xlsx", "excelDemo.xlsx")
		}
	}
}

//下载操作
func downFile(c *gin.Context, path string, fileName string) {
	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(path)
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
