package routers

import (
	"github.com/gin-gonic/gin"
	menuApi "rbac_manager/api/menu"
	"rbac_manager/middleware"
)

func MenuRouter(r *gin.RouterGroup) {
	g := r.Group("")
	g.Use(middleware.AuthMiddleware)
	menu := new(menuApi.Menu)
	g.POST("/menu", middleware.BindJson[menuApi.CreateMenuReq], menu.CreateMenu)
}
