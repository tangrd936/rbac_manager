package user

type UserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserLoginResp struct {
	Token string `json:"token"`
}
