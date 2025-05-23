package user

type UserLoginReq struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaId   string `json:"captcha_id"`
	CaptchaCode string `json:"captcha_code"`
}
type UserLoginResp struct {
	Token string `json:"token"`
}
type RegisterReq struct {
	EmailId    string `json:"email_id" binding:"required"`
	EmailCode  string `json:"email_code" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}
type RegisterResp struct {
}

type UpdatePwdReq struct {
	OldPwd   string `json:"old_pwd" binding:"required"`
	NewPwd   string `json:"new_pwd" binding:"required,max=64"`
	ReNewPwd string `json:"re_new_pwd" binding:"required,max=64"`
}
type UpdatePwdResp struct{}
