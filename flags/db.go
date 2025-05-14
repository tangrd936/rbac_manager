package flags

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"rbac_manager/global"
	"rbac_manager/models"
)

func AutoMigrate() {
	err := global.Db.AutoMigrate(
		&models.UserModel{},
		&models.RoleModel{},
		&models.UserRoleModel{},
		&models.MenuModel{},
		&models.ApiModel{},
		&models.RoleMenuModel{},
		&gormadapter.CasbinRule{},
	)
	if err != nil {
		global.Log.Error(fmt.Sprintf("auto migrate err: %v", err))
		return
	}
	global.Log.Info(fmt.Sprintf("auto migrate success"))
}
