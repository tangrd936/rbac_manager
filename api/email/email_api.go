package email

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"rbac_manager/common"
	"rbac_manager/global"
	"rbac_manager/middleware"
	capt "rbac_manager/utils/captcha"
	"rbac_manager/utils/email"
)

type Email struct{}

func (e *Email) SendEmail(c *gin.Context) {
	cr := middleware.GetReqData[SendEmailReq](c)
	if !global.Conf.Email.Enable {
		common.FailWithMsg(c, "管理员未配置邮箱,暂时无法注册")
		return
	}
	if !global.Conf.Captcha.Enable {
		common.FailWithMsg(c, "未启用验证码")
		return
	}
	var driver = base64Captcha.DriverString{
		Width:           200,
		Height:          60,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "0123456789",
	}

	cp := base64Captcha.NewCaptcha(&driver, capt.CaptchaStore)
	id, _, code, err := cp.Generate()
	if err != nil {
		global.Log.Error("generate base64 captcha failed", zap.Error(err))
		common.FailWithMsg(c, "generate captcha failed")
	}
	content := fmt.Sprintf("您正在完成用户注册，这是您的验证码\n %s \n验证吗5分钟内有效", code)
	mails := []string{cr.Email}
	err = email.SendMail(mails, "用户注册", content)
	if err != nil {
		global.Log.Error("send email failed", zap.Error(err))
		common.FailWithMsg(c, "send email failed")
	}

	common.OkWithData(c, SendEmailResp{
		EmailId: id,
	})

}
