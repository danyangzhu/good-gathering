package util

import (
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"good_gathering/conf"
)

var Region string

func Init() {
	rand.Seed(time.Now().UnixNano())
	initRegion()
}

func GetRegion() string {
	return Region
}

//获取地域信息
func initRegion() {
	var fileName = "/etc/region"
	regionConf := conf.Use("app")
	multiRegion := regionConf.MustBool("multiRegion", false)
	region := regionConf.MustString("region", "default")
	//非多地域模式，使用default配置
	if multiRegion == false {
		Region = "default"
		return
	}

	fd, err := os.Open(fileName)
	if err != nil {
		Region = region
		return
	}
	defer fd.Close()
	b, err := ioutil.ReadAll(fd)
	if err != nil {
		Region = region
		return
	}
	Region = string(b)
}
