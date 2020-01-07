package service

import (
	"ClockInLite/model"
	"ClockInLite/package/error"
	"encoding/json"
	"fmt"
)

type MenuId struct {
	ID int `form:"id" json:"id" binding:"required"`
}

type TopMenuInfo struct {
	Name    string `form:"name" json:"name" binding:"required"`
	OrderBy int    `form:"order_by" json:"order_by" binding:"required"`
}

type SubMenuInfo struct {
	ParentId        int    `form:"p_id" json:"p_id" binding:"required"`
	TopMenuInfo
	MenuRouter string `form:"menu_router" json:"menu_router" binding:"required"`
}

type TopMenu struct {
	MenuId
	TopMenuInfo
}

type SubMenu struct {
	MenuId
	SubMenuInfo
}

//添加顶级菜单
func (menuInfo *TopMenuInfo) TopMenuAdd() int {
	whereMaps := map[string]interface{}{"name": menuInfo.Name}
	isExist := model.ExistMenu(whereMaps)
	if isExist == true {
		return error.ERROR_EXIST_MENU
	}
	menu := model.Menu{
		Name:    menuInfo.Name,
		OrderBy: menuInfo.OrderBy,
	}
	if err := model.AddMenu(&menu); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//添加子菜单
func (menuInfo *SubMenuInfo) SubMenuAdd() int {
	whereMaps := map[string]interface{}{"name": menuInfo.Name}
	isExist := model.ExistMenu(whereMaps)
	if isExist == true {
		return error.ERROR_EXIST_MENU
	}

	menu := model.Menu{
		ParentId:        menuInfo.ParentId,
		Name:       menuInfo.Name,
		MenuRouter: menuInfo.MenuRouter,
		OrderBy:    menuInfo.OrderBy,
	}
	if err := model.AddMenu(&menu); err != nil {
		fmt.Println(err)
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//删除菜单
func (menuInfo *MenuId) MenuDel() int {
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

//编辑菜单
func (menuInfo *MenuId) MenuEdit() (model.Menu, int) {
	menu, err := model.GetMenu(map[string]interface{}{"id": menuInfo.ID})
	if err != nil {
		return menu, error.ERROR_NOT_EXIST_MENU
	}
	return menu, error.SUCCESS
}

//保存顶级菜单
func (TopMenu *TopMenu) TopMenuSave() int {
	id := TopMenu.ID
	menu := model.Menu{
		Name:    TopMenu.Name,
		OrderBy: TopMenu.OrderBy,
	}
	if err := model.SaveMenu(id, menu); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}

//保存子菜单
func (menuInfo *SubMenu) SubMenuSave() int {
	id := menuInfo.ID
	menu := model.Menu{
		ParentId:        menuInfo.ParentId,
		Name:       menuInfo.Name,
		MenuRouter: menuInfo.MenuRouter,
		OrderBy:    menuInfo.OrderBy,
	}
	if err := model.SaveMenu(id, menu); err != nil {
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
func GetTreeMenus() string {
	var menu model.Menu
	list := menu.GetTreeMenus(0)

	body, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
