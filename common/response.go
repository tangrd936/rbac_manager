package common

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rbac_manager/global"
	"rbac_manager/utils/validata"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func response(c *gin.Context, code int, data any, msg string) {
	r := Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	c.JSON(0, r)
}

func Ok(c *gin.Context, data any, msg string) {
	response(c, 0, data, msg)
}

func OkWithData(c *gin.Context, data any) {
	response(c, 0, data, "OK")
}

func OkWithMsg(c *gin.Context, msg string) {
	response(c, 0, gin.H{}, msg)
}

func Fail(c *gin.Context, code int, msg string, err error) {
	global.Log.Error(msg, zap.Error(err))
	response(c, code, gin.H{}, msg)
}

func FailWithMsg(c *gin.Context, msg string, err error) {
	global.Log.Error(msg, zap.Error(err))
	response(c, 1001, gin.H{}, msg)
}

func FailWithError(c *gin.Context, err error) {
	global.Log.Error(err.Error())
	response(c, 1001, gin.H{}, err.Error())
}

func FailWithBindParam(c *gin.Context, err error) {
	global.Log.Error(err.Error())
	resp := validata.ValidateErr(err)
	response(c, 1001, resp.Field, resp.Msg)
}
