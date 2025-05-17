package routers

import (
	"github.com/gin-gonic/gin"
	captchaApi "rbac_manager/api/captcha"
)

func CaptchaRouter(r *gin.RouterGroup) {
	g := r.Group("/captcha")
	capt := new(captchaApi.Captcha)
	g.GET("/generate", capt.GenerateCaptcha)
}
