package user

import "rbac_manager/models"

// 用户登录
type UserLoginReq struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaId   string `json:"captcha_id"`
	CaptchaCode string `json:"captcha_code"`
}
type UserLoginResp struct {
	Token string `json:"token"`
}

// 用户注册
type RegisterReq struct {
	EmailId    string `json:"email_id" binding:"required"`
	EmailCode  string `json:"email_code" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}
type RegisterResp struct {
}

// 修改密码
type UpdatePwdReq struct {
	OldPwd   string `json:"old_pwd" binding:"required"`
	NewPwd   string `json:"new_pwd" binding:"required,max=64"`
	ReNewPwd string `json:"re_new_pwd" binding:"required,max=64"`
}
type UpdatePwdResp struct{}

// 修改用户信息
type UpdateUserInfoReq struct {
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}
type UpdateUserInfoResp struct{}

// 获取用户信息
type GetUserInfoReq struct {
}
type GetUserInfoResp struct {
	UserId   uint   `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	RoleList []uint `json:"role_list"`
}

// 查询用户信息列表
type PageInfo struct {
	Limit int    `form:"limit"` // 每页展示条数
	Page  int    `form:"page"`  // 当前页
	Sort  uint   `form:"sort"`  // 排序
	Key   string `form:"key"`   // 模糊匹配
}
type GetUserInfoListReq struct {
	PageInfo
	Role     uint   `form:"role"`
	UserName string `form:"user_name"`
	Email    string `form:"email"`
}
type GetUserInfoListResp struct {
	UserInfoList []models.UserModel `json:"user_info_list"`
	Count        int                `json:"count"`
}

// 删除用户
type DeleteUserReq struct {
	IdList []uint `json:"id_list"`
}
type DeleteUserResp struct {
}
