package service

import (
	"ClockInLite/model"
	"ClockInLite/package/error"
)

type MenuId struct {
	ID int `form:"id" json:"id" binding:"required"`
}

type MenuInfo struct {
	ParentId   int    `form:"parent_id" json:"parent_id"`
	Name       string `form:"name" json:"name" binding:"required"`
	MenuRouter string `form:"menu_router" json:"menu_router" binding:"required"`
	OrderBy    int    `form:"order_by" json:"order_by"`
}

type Menu struct {
	MenuId
	MenuInfo
}

//添加菜单
func (menuInfo *MenuInfo) AddMenu() int {
	whereMaps := map[string]interface{}{"name": menuInfo.Name}
	isExist := model.ExistMenu(whereMaps)
	if isExist == true {
		return error.ERROR_EXIST_MENU
	}
	menu := model.Menu{
		ParentId:   menuInfo.ParentId,
		Name:       menuInfo.Name,
		OrderBy:    menuInfo.OrderBy,
		MenuRouter: menuInfo.MenuRouter,
	}
	if err := model.AddMenu(&menu); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//删除菜单
func (menuInfo *MenuId) DelMenu() int {
	id := map[string]interface{}{"id": menuInfo.ID}
	isExist := model.ExistMenu(id)
	if isExist == false {
		return error.ERROR_NOT_EXIST_MENU
	}

	err := model.DelMenu(map[string]interface{}{"id": menuInfo.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}
	return error.SUCCESS
}

//获取菜单
func (menuInfo *MenuId) GetMenu() (model.Menu, int) {
	menu, err := model.GetMenu(map[string]interface{}{"id": menuInfo.ID})
	if err != nil {
		return menu, error.ERROR_NOT_EXIST_MENU
	}
	return menu, error.SUCCESS
}

//保存顶级菜单
func (menu *Menu) SaveMenu() int {
	id := menu.ID

	menuInfo := model.Menu{
		ParentId:   menu.ParentId,
		Name:       menu.Name,
		MenuRouter: menu.MenuRouter,
		OrderBy:    menu.OrderBy,
	}
	if err := model.SaveMenu(id, menuInfo); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}

//获取用户菜单
func GetAdminMenus(adminId int) (list []model.Menu) {
	list, _ = model.GetMenuByAdminId(adminId)
	return
}

//获取菜单树
//func GetTreeMenus() string {
//	var menu model.Menu
//	list := menu.GetTreeMenus(0)
//	//fmt.Println("list:", list)
//	//
//	//body, err := json.Marshal(list)
//	//fmt.Println("body:", body)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//return string(body)
//	return list
//}
