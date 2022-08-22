package service

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// GetAllPrice 获取所有价格
func GetAllPrice() ([]string,error) {
	url := "http://www.kyzb9999.cn/price_get_rtj_all.php?t=1659889904555"
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	// 切割数据
	countSplit := strings.Split(string(body), ",")
	priceSlice := countSplit[1 : len(countSplit)-5]

	return priceSlice,nil
}

