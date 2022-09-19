package main

import (
	"flag"
	"good_gathering/conf"
	"good_gathering/controller"
	"good_gathering/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
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
		api.GET("/DescribeAllPrice", controller.DescribeAllPrice)
		// 用于描述首页 所有价格
		api.POST("/DescribeAllPrice", controller.DescribeAllPrice)
		// 用于描述行情页数据
		api.POST("/DescribePriceData", controller.DescribePriceData)
	}

	router.Use(TlsHandler(443))
	router.RunTLS(":"+strconv.Itoa(443), "./conf/kyzb0755.com_bundle.pem", "./conf/kyzb0755.com.key")

}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
