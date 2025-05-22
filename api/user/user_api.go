package user

import (
	"github.com/gin-gonic/gin"
	"rbac_manager/common"
	"rbac_manager/global"
	"rbac_manager/middleware"
	"rbac_manager/models"
	capt "rbac_manager/utils/captcha"
	"rbac_manager/utils/jwts"
	"rbac_manager/utils/pwd"
)

type User struct {
}

func (u *User) Login(c *gin.Context) {
	cr := middleware.GetReqData[UserLoginReq](c)
	if global.Conf.Captcha.Enable {
		if cr.CaptchaId == "" || cr.CaptchaCode == "" {
			common.FailWithMsg(c, "请输入图片验证码", nil)
			return
		}
		if !capt.CaptchaStore.Verify(cr.CaptchaId, cr.CaptchaCode, true) {
			common.FailWithMsg(c, "图片验证码验证失败", nil)
			return
		}
	}
	var user models.UserModel
	err := global.Db.Preload("RoleList").Take(&user, "user_name = ?", cr.Username).Error
	if err != nil {
		common.FailWithMsg(c, "用户名或密码错误", err)
		return
	}
	if pwd.ComparePasswords(user.Password, cr.Password) {
		common.FailWithMsg(c, "用户名或密码错误", nil)
		return
	}
	var roleList []uint
	for _, role := range user.RoleList {
		roleList = append(roleList, role.ID)
	}

	token, err := jwts.GetToken(jwts.ClaimUserInfo{
		UserId:   user.ID,
		UserName: user.UserName,
		RoleList: roleList,
	})
	if err != nil {
		common.FailWithMsg(c, "用户登陆失败", err)
		return
	}
	res := UserLoginResp{Token: token}
	common.Ok(c, res, "用户登录成功")
}

func (u *User) Register(c *gin.Context) {
	cr := middleware.GetReqData[RegisterReq](c)

	// 验证发送验证码的邮箱是否与注册邮箱一致
	val := global.Redis.Get(c.Request.Context(), cr.Email).Val()
	if val == "" || val != cr.EmailCode {
		common.FailWithMsg(c, "验证码或邮箱不匹配", nil)
		return
	}

	// 验证邮箱
	if !capt.CaptchaStore.Verify(cr.EmailId, cr.EmailCode, true) {
		common.FailWithMsg(c, "邮箱验证失败", nil)
		return
	}
	// 判断两次密码是否一直
	if cr.Password != cr.RePassword {
		common.FailWithMsg(c, "两次密码输入不一致", nil)
		return
	}
	// 判断邮箱是否使用
	var user models.UserModel
	err := global.Db.Take(&user, "email = ?", cr.Email).Error
	if err == nil {
		common.FailWithMsg(c, "邮箱已被使用", err)
		return
	}
	//创建用户
	password := pwd.HashPassword(cr.Password)
	err = global.Db.Create(&models.UserModel{
		UserName: cr.Email,
		Email:    cr.Email,
		Password: password,
	}).Error
	if err != nil {
		common.FailWithMsg(c, "用户注册失败", err)
		return
	}
	// 删除redis中的email key
	global.Redis.Del(c.Request.Context(), cr.Email)

	common.OkWithMsg(c, "用户注册成功")
}
