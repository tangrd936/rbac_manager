package routers

import (
	"github.com/gin-gonic/gin"
	emailApi "rbac_manager/api/email"
	"rbac_manager/middleware"
)

func EmailRouter(r *gin.RouterGroup) {
	g := r.Group("/email")
	mail := new(emailApi.Email)
	g.POST("/send", middleware.BindJson[emailApi.SendEmailReq], mail.SendEmail)
}
