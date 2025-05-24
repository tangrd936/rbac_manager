package routers

import (
	"github.com/gin-gonic/gin"
	imageApi "rbac_manager/api/image"
	"rbac_manager/middleware"
)

func ImageRouter(r *gin.RouterGroup) {
	g := r.Group("/image")
	g.Use(middleware.AuthMiddleware)
	image := new(imageApi.Image)
	g.POST("/upload", image.UploadImage)
}
