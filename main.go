package main

import (
	"flag"
	"good_gathering/conf"
	"good_gathering/controller"
	"good_gathering/service"

	"github.com/gin-gonic/gin"
)

// SetCmdParam 设置参数
func SetCmdParam() {
	var confDir string
	flag.StringVar(&confDir, "confDir", "./conf", "config dir")
	flag.Parse()
	switch {
	case confDir != "./conf":
		conf.SetConfDir(confDir)
	default:
	}
}

func main() {
	SetCmdParam()
	go service.TaskInit()
	router := gin.Default()

	api := router.Group("/api")
	{
		// 用于描述首页 所有价格
		api.POST("/DescribeAllPrice", controller.DescribeAllPrice)
		// 用于描述行情页数据
		api.POST("/DescribePriceData", controller.DescribePriceData)
	}

	router.Run(":80")
}
