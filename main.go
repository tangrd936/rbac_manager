package main

import (
	"rbac_manager/core"
	"rbac_manager/flags"
	"rbac_manager/routers"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	core.InitLogger("logs") // 初始化日志配置
	core.InitConfig()       // 初始化系统配置
	core.InitDB()           // 加载数据库
	core.InitRedis()        // 加载redis
	core.InitCasbin()       // 加载casbin

	// 命令行参数运行
	flags.Run()
	// 运行web服务
	routers.Run()
}
