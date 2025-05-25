package middleware

import (
	"github.com/gin-gonic/gin"
	"rbac_manager/common"
)

func BindJson[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		common.FailWithBindParam(c, err)
		c.Abort() // 直接拦截返回
		return
	}
	c.Set("req", cr)
	c.Next()
	return
}

func BindQuery[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		common.FailWithBindParam(c, err)
		c.Abort() // 直接拦截返回
		return
	}
	c.Set("req", cr)
	c.Next()
	return
}

func GetReqData[T any](c *gin.Context) (data T) {
	return c.MustGet("req").(T)
}
