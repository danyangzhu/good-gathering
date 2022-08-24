package service

import (
	"fmt"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

const Measurement = "price"
const TickerTime = 60

var infoNameList = []string{"黄金", "白银", "铂金", "钯金", "黄金99.99", "黄金(T+D)", "白银(T+D)", "铂金99.95",
	"美黄金", "美铂金", "美钯金", "美白银", "美铑金", "港金", "伦敦金", "伦敦银", "伦敦铂", "伦敦钯", "美元"}

func WriteDB() {
	allPrice, err := GetAllPrice()
	if err != nil {

	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://127.0.0.1:8086",
	})
	if err != nil {
		fmt.Println(err)
	}

	for i, name := range infoNameList {
		tags := map[string]string{
			"price_type": name,
		}

		fields := map[string]interface{}{
			"repo_price":  StringToFloat64(allPrice[4*i]),
			"shell_price": StringToFloat64(allPrice[4*i+1]),
			"max_price":   StringToFloat64(allPrice[4*i+2]),
			"min_price":   StringToFloat64(allPrice[4*i+3]),
		}
		WriteInfluxdb(c, Measurement, tags, fields)
	}
}

func StringToFloat64(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	return float
}

func TaskInit() {
	WriteDB()
	ticker := time.Tick(TickerTime * time.Second)
	for {
		<-ticker
		WriteDB()
	}
}
