package controller

import (
	"good_gathering/service"

	"github.com/gin-gonic/gin"
)

func DescribeAllPrice(c *gin.Context) {

	allPrice, err := service.GetAllPrice()
	if err != nil {

	}

	c.JSON(200, gin.H{
		"黄金回购":      allPrice[0],
		"黄金销售":      allPrice[1],
		"黄金最高价":     allPrice[2],
		"黄金最低价":     allPrice[3],
		"白银回购":      allPrice[4],
		"白银销售":      allPrice[5],
		"白银最高价":     allPrice[6],
		"白银最低价":     allPrice[7],
		"铂金回购":      allPrice[8],
		"铂金销售":      allPrice[9],
		"铂金最高价":     allPrice[10],
		"铂金最低价":     allPrice[11],
		"钯金回购":      allPrice[8],
		"钯金销售":      allPrice[9],
		"钯金最高价":     allPrice[10],
		"钯金最低价":     allPrice[11],
		"黄金9999回购":  allPrice[12],
		"黄金9999销售":  allPrice[13],
		"黄金9999最高价": allPrice[14],
		"黄金9999最低价": allPrice[15],
		"黄金T+D回购":   allPrice[16],
		"黄金T+D销售":   allPrice[17],
		"黄金T+D最高价":  allPrice[18],
		"黄金T+D最低价":  allPrice[19],
		"白银T+D回购":   allPrice[20],
		"白银T+D销售":   allPrice[21],
		"白银T+D最高价":  allPrice[22],
		"白银T+D最低价":  allPrice[23],
		"铂金9995回购":  allPrice[24],
		"铂金9995销售":  allPrice[25],
		"铂金9995最高价": allPrice[26],
		"铂金9996最低价": allPrice[27],
	})
}

func DescribePriceData(c *gin.Context) {

}
