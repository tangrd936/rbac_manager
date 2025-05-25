package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rbac_manager/global"
	"strconv"
)

func Run() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	g := r.Group("/api")
	UserRouter(g)
	CaptchaRouter(g)
	EmailRouter(g)
	ImageRouter(g)
	MenuRouter(g)

	//配置静态资源访问
	r.Static("/static", "static")

	addr := global.Conf.System.IP + ":" + strconv.Itoa(global.Conf.System.Port)
	err := r.Run(addr)
	if err != nil {
		global.Log.Error(fmt.Sprintf("listen %s fail", addr))
		return
	}
}
