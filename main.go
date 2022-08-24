package main

import (
	"good_gathering/controller"
	"good_gathering/service"

	"github.com/gin-gonic/gin"
)

func main() {
	go service.TaskInit()
	router := gin.Default()

	api := router.Group("/api")
	{
		// 用于描述首页 所有价格
		api.POST("/DescribeAllPrice", controller.DescribeAllPrice)
		// 用于描述行情页数据
		api.POST("/DescribePriceData", controller.DescribePriceData)
	}

	router.Run(":8080")
}
