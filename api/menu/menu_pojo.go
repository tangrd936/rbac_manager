package menu

// 添加菜单
type CreateMenuReq struct {
	Icon         string `json:"icon"`
	Title        string `json:"title" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Path         string `json:"path" binding:"required"`
	Component    string `json:"component"`
	ParentMenuId *uint  `json:"parent_menu_id"`
	Sort         int    `json:"sort"`
}

type CreateMenuResp struct{}
