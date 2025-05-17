package captcha

type GenerateCaptchaReq struct {
}
type GenerateCaptchaResp struct {
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}
