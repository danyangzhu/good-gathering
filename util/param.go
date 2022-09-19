package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type BaseParam struct {
	Action        string `form:"Action"   binding:"required"` //执行函数名
	Module        string `form:"Module" binding:"required"`   //模块名，即路由的group组名
	AppId         int    `form:"AppId"   binding:"exists"`
	Uin           string `form:"Uin"   binding:"required"`
	SubAccountUin string `form:"SubAccountUin"   binding:"required"`
	Region        string `form:"Region"   binding:"required"`
	RequestId     string `form:"RequestId"`
	Version       string `form:"Version"`
	RequestSource string `form:"RequestSource"`
	ClientIp      string `form:"ClientIp"`
	Language      string `form:"Language"`
}

func GetRequestIdFromBody(c *gin.Context) (string, error) {
	var param BaseParam
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return "", err
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = json.Unmarshal(body, &param)
	if err != nil {
		return "", err
	}

	return param.RequestId, nil
}

// GetRequestId 获取requestid，如果不存在，则随机生成
func GetRequestId(c *gin.Context) string {
	if rid := c.Query("RequestId"); rid != "" {
		return rid
	}

	if rid, ok := c.Get("RequestId"); ok {
		return rid.(string)
	}

	if rid, err := GetRequestIdFromBody(c); err == nil {
		c.Set("RequestId", rid)
		return rid
	}

	rid := RandString(32)
	c.Set("RequestId", rid)
	return rid
}
