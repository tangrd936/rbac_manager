package user

import (
	"fmt"
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

// Login 用户登录
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
	fmt.Printf("密码：%s\n", cr.Password)
	var user models.UserModel
	err := global.Db.Preload("RoleList").Take(&user, "user_name = ?", cr.Username).Error
	if err != nil {
		common.FailWithMsg(c, "用户名或密码错误", err)
		return
	}
	if !pwd.ComparePasswords(user.Password, cr.Password) {
		common.FailWithMsg(c, "用户名或密码错误", nil)
		return
	}

	token, err := jwts.GetToken(jwts.ClaimUserInfo{
		UserId:   user.ID,
		UserName: user.UserName,
		RoleList: user.GetRoleList(),
	})
	if err != nil {
		common.FailWithMsg(c, "用户登陆失败", err)
		return
	}
	res := UserLoginResp{Token: token}
	common.Ok(c, res, "用户登录成功")
}

// Register 用户注册
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

// UpdatePassword 修改密码
func (u *User) UpdatePassword(c *gin.Context) {
	cr := middleware.GetReqData[UpdatePwdReq](c)
	auth := middleware.GetAuth(c)
	global.Log.Info(fmt.Sprintf("用户信息：%s", auth))
	var user models.UserModel
	err := global.Db.Take(&user, auth.UserId).Error
	if err != nil {
		common.FailWithMsg(c, "用户不存在", err)
		return
	}
	if !pwd.ComparePasswords(user.Password, cr.OldPwd) {
		common.FailWithMsg(c, "原密码错误", nil)
		return
	}

	if cr.NewPwd != cr.ReNewPwd {
		common.FailWithMsg(c, "两次输入密码不一致", nil)
		return
	}
	// 新密码加密
	newPwd := pwd.HashPassword(cr.NewPwd)
	err = global.Db.Model(&user).Update("password", newPwd).Error
	if err != nil {
		common.FailWithMsg(c, "修改密码失败", err)
		return
	}

	common.OkWithMsg(c, "修改密码成功")
}

// UpdateUserInfo 更新用户信息
func (u *User) UpdateUserInfo(c *gin.Context) {
	cr := middleware.GetReqData[UpdateUserInfoReq](c)
	auth := middleware.GetAuth(c)
	global.Log.Info(fmt.Sprintf("用户信息：%s", auth))
	var user models.UserModel
	err := global.Db.Take(&user, auth.UserId).Error
	if err != nil {
		common.FailWithMsg(c, "用户不存在", err)
		return
	}

	err = global.Db.Model(&user).Updates(models.UserModel{
		NickName: cr.NickName,
		Avatar:   cr.AvatarUrl,
	}).Error
	if err != nil {
		common.FailWithMsg(c, "更新用户信息失败", err)
		return
	}
	common.OkWithMsg(c, "更新用户信息成功")
}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo(c *gin.Context) {
	auth := middleware.GetAuth(c)
	global.Log.Info(fmt.Sprintf("用户信息：%s", auth))
	var user models.UserModel
	err := global.Db.Preload("RoleList").Take(&user, auth.UserId).Error
	if err != nil {
		common.FailWithMsg(c, "用户不存在", err)
		return
	}

	data := GetUserInfoResp{
		UserId:   user.ID,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		RoleList: user.GetRoleList(),
	}
	common.OkWithData(c, data)
}

// GetUserInfoList 获取用户信息列表
func (u *User) GetUserInfoList(c *gin.Context) {
	auth := middleware.GetAuth(c)
	global.Log.Info(fmt.Sprintf("登录用户信息：%+v", auth))
	var user models.UserModel
	err := global.Db.Preload("RoleList").Take(&user, auth.UserId).Error
	if err != nil {
		common.FailWithMsg(c, "用户不存在", err)
		return
	}
	// 判断当前用户是否为管理员
	if !user.IsAdmin {
		common.FailWithMsg(c, "当前用户不是管理员", err)
		return
	}

	cr := middleware.GetReqData[GetUserInfoListReq](c)
	list := make([]models.UserModel, 0)
	offset := (cr.Page - 1) * cr.Limit
	global.Db.Preload("RoleList").Where(models.UserModel{
		UserName: cr.UserName,
		Email:    cr.Email,
	}).Where("nick_name like ?", fmt.Sprintf("%%%s%%", cr.Key)).Limit(cr.Limit).Offset(offset).Order(cr.Sort).Find(&list)

	data := GetUserInfoListResp{
		UserInfoList: list,
		Count:        len(list),
	}
	common.OkWithData(c, data)
}

// DelUser 获取用户信息列表
func (u *User) DelUser(c *gin.Context) {
	auth := middleware.GetAuth(c)
	global.Log.Info(fmt.Sprintf("登录用户信息：%+v", auth))
	var user models.UserModel
	err := global.Db.Preload("RoleList").Take(&user, auth.UserId).Error
	if err != nil {
		common.FailWithMsg(c, "用户不存在", err)
		return
	}
	// 判断当前用户是否为管理员
	if !user.IsAdmin {
		common.FailWithMsg(c, "当前用户不是管理员", err)
		return
	}

	cr := middleware.GetReqData[DeleteUserReq](c)
	list := make([]models.UserModel, 0)
	global.Db.Find(&list, "id in ?", cr.IdList)
	var count int64
	if len(list) > 0 {
		count = global.Db.Delete(&list).RowsAffected
	}

	msg := fmt.Sprintf("删除用户 %d 个", count)
	common.OkWithMsg(c, msg)
}
