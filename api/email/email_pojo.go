package email

type SendEmailReq struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaId   string `json:"captcha_id"`
	CaptchaCode string `json:"captcha_code"`
}
type SendEmailResp struct {
	EmailId string `json:"email_id"`
}
