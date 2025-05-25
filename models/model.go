package models

import (
	"fmt"
	"gorm.io/gorm"
	"rbac_manager/global"
)

type UserModel struct {
	gorm.Model
	UserName string      `gorm:"size:32;unique" json:"user_name"`
	NickName string      `gorm:"size:32" json:"nick_name"`
	Avatar   string      `gorm:"size:256" json:"avatar"`
	Email    string      `gorm:"size:128" json:"email"`
	Password string      `gorm:"size:64" json:"-"`
	IsAdmin  bool        `gorm:"default:false" json:"is_admin"`
	RoleList []RoleModel `gorm:"many2many:user_role_models;joinForeignKey:UserId;joinReferences:RoleId" json:"role_list"`
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) error {
	var userRoleList []UserRoleModel
	err := tx.Find(&userRoleList, "user_id = ?", u.ID).Delete(&userRoleList).Error
	global.Log.Info(fmt.Sprintf("删除用户角色关联表 %d 条数据", len(userRoleList)))
	return err
}

func (u *UserModel) GetRoleList() []uint {
	var roleList []uint
	for _, role := range u.RoleList {
		roleList = append(roleList, role.ID)
	}
	return roleList
}

type RoleModel struct {
	gorm.Model
	Title    string      `gorm:"size:16;unique" json:"title"`
	UserList []UserModel `gorm:"many2many:user_role_models;joinForeignKey:RoleId;joinReferences:UserId" json:"user_list"`
	MenuList []MenuModel `gorm:"many2many:role_menu_models;joinForeignKey:RoleId;joinReferences:MenuId" json:"role_list"`
}

type UserRoleModel struct {
	gorm.Model
	UserId    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserId" json:"-"`
	RoleId    uint      `json:"role_id"`
	RoleModel RoleModel `gorm:"foreignKey:RoleId" json:"-"`
}

type Meta struct {
	Icon  string `gorm:"size:256" json:"icon"`
	Title string `gorm:"size:16" json:"title"`
}
type MenuModel struct {
	gorm.Model
	Name            string `gorm:"size:32;unique" json:"name"`
	Path            string `gorm:"size:64" json:"path"`
	Component       string `gorm:"size:128" json:"component"`
	Meta            `gorm:"embedded" json:"meta"`
	ParentMenuId    *uint        `json:"parent_menu_id"`
	ParentMenuModel *MenuModel   `gorm:"foreignKey:ParentMenuId" json:"-"`
	Children        []*MenuModel `gorm:"foreignKey:ParentMenuId" json:"children"`
	Sort            int          `json:"sort"`
}

type ApiModel struct {
	gorm.Model
	Name   string `gorm:"size:32;unique" json:"name"`
	Path   string `gorm:"size:64" json:"path"`
	Method string `gorm:"size:128" json:"method"`
	Group  string `gorm:"size:128" json:"group"`
}

type RoleMenuModel struct {
	gorm.Model
	RoleId    uint      `json:"role_id"`
	RoleModel RoleModel `gorm:"foreignKey:RoleId" json:"-"`
	MenuId    uint      `json:"menu_id"`
	MenuModel MenuModel `gorm:"foreignKey:MenuId" json:"-"`
}
