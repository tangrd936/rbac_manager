package core

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"rbac_manager/global"
)

// InitCasbin 初始化casbin
func InitCasbin() {
	adapter, err := gormadapter.NewAdapterByDB(global.Db)
	if err != nil {
		global.Log.Error(fmt.Sprintf("Casbin adapter creation failed:%v", err))
		return
	}
	m, err := model.NewModelFromFile("config/rbac_model.pml")
	if err != nil {
		global.Log.Error(fmt.Sprintf("Casbin load model failed:%v", err))
		return
	}

	// 加载RBAC模型配置
	enforcer, err := casbin.NewCachedEnforcer(m, adapter)
	if err != nil {
		global.Log.Error(fmt.Sprintf("Casbin NewCachedEnforcer failed:%v", err))
		return
	}
	enforcer.SetExpireTime(60 * 60)
	err = enforcer.LoadPolicy()
	if err != nil {
		global.Log.Error(fmt.Sprintf("Casbin load policy failed:%v", err))
		return
	}
	global.Casbin = enforcer
	global.Log.Info(fmt.Sprintf("Casbin initialized"))
}
