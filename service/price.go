package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/influxdata/influxdb1-client/models"

	client "github.com/influxdata/influxdb1-client/v2"
)

// GetAllPrice 获取所有价格
func GetAllPrice() ([]string, error) {
	url := "http://www.kyzb9999.cn/price_get_rtj_all.php?t=1659889904555"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// 切割数据
	countSplit := strings.Split(string(body), ",")
	priceSlice := countSplit[1 : len(countSplit)-5]

	return priceSlice, nil
}

func GetPriceData(priceType string) models.Row {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	startTime := GetNowDayFirstTime()
	sqlStr := fmt.Sprintf("SELECT max(\"shell_price\") FROM price where price_type = '%s' and time > %d group by time(1m)", priceType, startTime)
	q := client.NewQuery(sqlStr, Database, "rfc3339")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		if len(response.Results) > 0 && len(response.Results[0].Series) > 0 {
			return response.Results[0].Series[0]
		}
	}
	return models.Row{}
}

// GetNowDayFirstTime 获取当天0点的Unix时间戳 纳秒级别
func GetNowDayFirstTime() int64 {
	timeNow := time.Now()
	// 获取当天0点时间 time类型
	timeToday := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())

	return timeToday.UnixNano()
}
