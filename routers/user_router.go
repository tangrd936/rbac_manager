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
}
