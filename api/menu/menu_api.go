package menu

import (
	"github.com/gin-gonic/gin"
	"rbac_manager/common"
	"rbac_manager/global"
	"rbac_manager/middleware"
	"rbac_manager/models"
)

type Menu struct {
}

func (m *Menu) CreateMenu(c *gin.Context) {
	cr := middleware.GetReqData[CreateMenuReq](c)
	var menu models.MenuModel
	err := global.Db.Take(&menu, "name = ?", cr.Name).Error
	if err == nil {
		common.FailWithMsg(c, "菜单名不能重复", err)
		return
	}
	if cr.ParentMenuId != nil {
		var parentMenu models.MenuModel
		err = global.Db.Take(&parentMenu, "id = ?", *cr.ParentMenuId).Error
		if err != nil {
			common.FailWithMsg(c, "父菜单不存在", err)
			return
		}
	}
	menuData := models.MenuModel{
		Name:      cr.Name,
		Path:      cr.Path,
		Component: cr.Component,
		Meta: models.Meta{
			Icon:  cr.Icon,
			Title: cr.Title,
		},
		ParentMenuId: cr.ParentMenuId,
		Sort:         cr.Sort,
	}
	err = global.Db.Create(&menuData).Error
	if err != nil {
		common.FailWithMsg(c, "添加菜单失败", err)
	}
	common.OkWithMsg(c, "菜单创建成功")
}
