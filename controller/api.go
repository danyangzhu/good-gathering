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

// DescribeAllPrice 查询所有价格
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
		infos = append(infos, PriceInfo{
			Name:      name,
			RepoPrice: allPrice[4*i],
			SellPrice: allPrice[4*i+1],
			MaxPrice:  allPrice[4*i+2],
			MinPrice:  allPrice[4*i+3],
		})
	}

	c.JSON(200, gin.H{
		"Response": infos,
	})
}

// DescribePriceData 查询当日的价格曲线
func DescribePriceData(c *gin.Context) {
	type Param struct {
		PriceType string `binding:"required"` // 黄金类型
	}

	var param Param
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	priceData := service.GetPriceData(param.PriceType)
	c.JSON(200, gin.H{
		"Response": priceData,
	})
}
