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
			common.FailWithMsg(c, "请输入图片验证码")
			return
		}
		if capt.CaptchaStore.Verify(cr.CaptchaId, cr.CaptchaCode, true) {
			common.OkWithMsg(c, "图片验证码验证失败")
			return
		}
	}
	var user models.UserModel
	err := global.Db.Preload("RoleList").Take(&user, "user_name = ?", cr.Username).Error
	if err != nil {
		common.FailWithMsg(c, "用户名或密码错误")
		return
	}
	if pwd.ComparePasswords(user.Password, cr.Password) {
		common.FailWithMsg(c, "用户名或密码错误")
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
		global.Log.Error("jwt颁发token失败：" + err.Error())
		common.FailWithMsg(c, "用户登陆失败")
		return
	}
	res := UserLoginResp{Token: token}
	common.Ok(c, res, "用户登录成功")
}
