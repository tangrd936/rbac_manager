package routers

import (
	"github.com/gin-gonic/gin"
	userApi "rbac_manager/api/user"
	"rbac_manager/middleware"
)

func UserRouter(r *gin.RouterGroup) {
	g := r.Group("/user")
	user := new(userApi.User)
	g.POST("/login", middleware.BindJson[userApi.UserLoginReq], user.Login)
	g.POST("/register", middleware.BindJson[userApi.RegisterReq], user.Register)
	g.PUT("/password", middleware.AuthMiddleware, middleware.BindJson[userApi.UpdatePwdReq], user.UpdatePassword)
	g.PUT("/info", middleware.AuthMiddleware, middleware.BindJson[userApi.UpdateUserInfoReq], user.UpdateUserInfo)
	g.GET("/info", middleware.AuthMiddleware, user.GetUserInfo)
	g.GET("/list", middleware.AuthMiddleware, middleware.BindQuery[userApi.GetUserInfoListReq], user.GetUserInfoList)
	g.DELETE("/info", middleware.AuthMiddleware, middleware.BindJson[userApi.DeleteUserReq], user.DelUser)
}
