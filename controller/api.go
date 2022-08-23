package controller

import (
	"good_gathering/service"

	"github.com/gin-gonic/gin"
)

type PriceInfo struct {
	Name      string
	RepoPrice string
	SellPrice string
	MaxPrice  string
	MinPrice  string
}

var infoNameList = []string{"黄金", "白银", "铂金", "钯金", "黄金99.99", "黄金(T+D)", "白银(T+D)", "铂金99.95",
	"美黄金", "美铂金", "美钯金", "美白银", "美铑金", "港金", "伦敦金", "伦敦银", "伦敦铂", "伦敦钯", "美元"}

func DescribeAllPrice(c *gin.Context) {
	allPrice, err := service.GetAllPrice()
	if err != nil {

	}
	var infos []PriceInfo
	for i, name := range infoNameList {
		if 4*(i+1) > len(allPrice) {
			// error log
			break
		}
		var info PriceInfo
		info.Name = name
		info.RepoPrice = allPrice[i]
		info.SellPrice = allPrice[i+1]
		info.MaxPrice = allPrice[i+2]
		info.MinPrice = allPrice[i+3]
		infos = append(infos, info)
	}

	c.JSON(200, gin.H{"response": infos})
}

func DescribePriceData(c *gin.Context) {

}
