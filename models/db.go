package models

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"rbac_manager/global"
)

func AutoMigrate() {
	err := global.Db.AutoMigrate(
		&UserModel{},
		&RoleModel{},
		&UserRoleModel{},
		&MenuModel{},
		&ApiModel{},
		&RoleMenuModel{},
		&gormadapter.CasbinRule{},
	)
	if err != nil {
		global.Log.Error(fmt.Sprintf("auto migrate err: %v", err))
		return
	}
	global.Log.Info(fmt.Sprintf("auto migrate success"))
}
