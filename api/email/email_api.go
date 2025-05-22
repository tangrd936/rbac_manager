package email

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"rbac_manager/common"
	"rbac_manager/global"
	"rbac_manager/middleware"
	capt "rbac_manager/utils/captcha"
	"rbac_manager/utils/email"
	"time"
)

type Email struct{}

func (e *Email) SendEmail(c *gin.Context) {
	cr := middleware.GetReqData[SendEmailReq](c)
	if !global.Conf.Email.Enable {
		common.FailWithMsg(c, "管理员未配置邮箱,暂时无法注册", nil)
		return
	}
	if !global.Conf.Captcha.Enable {
		common.FailWithMsg(c, "未启用验证码", nil)
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
		common.FailWithMsg(c, "generate captcha failed", err)
		return
	}
	content := fmt.Sprintf("您正在完成用户注册，这是您的验证码\n %s \n验证吗5分钟内有效", code)
	//设置当前邮箱验证码5分钟有效期
	err = global.Redis.Set(c.Request.Context(), cr.Email, code, 5*time.Minute).Err()
	if err != nil {
		common.FailWithMsg(c, "redis set base64 captcha failed", err)
		return
	}
	mails := []string{cr.Email}
	err = email.SendMail(mails, "用户注册", content)
	if err != nil {
		common.FailWithMsg(c, "send email failed", err)
		return
	}

	common.OkWithData(c, SendEmailResp{
		EmailId: id,
	})

}
