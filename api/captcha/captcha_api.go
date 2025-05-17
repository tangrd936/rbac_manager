package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"rbac_manager/common"
	"rbac_manager/global"
	capt "rbac_manager/utils/captcha"
)

type Captcha struct{}

func (captcha Captcha) GenerateCaptcha(c *gin.Context) {
	var driver = base64Captcha.DriverString{
		Width:           200,
		Height:          60,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "0123456789",
	}

	cp := base64Captcha.NewCaptcha(&driver, capt.CaptchaStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.Log.Error("generate base64 captcha failed", zap.Error(err))
		common.FailWithMsg(c, "generate captcha failed")
	}

	common.OkWithData(c, GenerateCaptchaResp{
		CaptchaID: id,
		Captcha:   b64s,
	})
}
