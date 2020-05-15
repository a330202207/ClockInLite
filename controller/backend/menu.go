package backend

import (
	"ClockInLite/config"
	"ClockInLite/model"
	"ClockInLite/package/error"
	"ClockInLite/service"
	"ClockInLite/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//菜单列表页
func GetMenuList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.ServerSetting.PageSize)))

	Offset := (page - 1) * pageSize

	Menus, count, _ := model.GetMenuList(pageSize, Offset, "id asc")

	util.JsonSuccessPage(c, count, Menus)
}

//添加顶级菜单
func TopMenuAdd(c *gin.Context) {
	var menu service.TopMenuInfo
	if err := c.ShouldBindJSON(&menu); err == nil {
		resCode := menu.TopMenuAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//添加子菜单
func SubMenuAdd(c *gin.Context) {
	var menu service.SubMenuInfo
	if err := c.ShouldBindJSON(&menu); err == nil {
		resCode := menu.SubMenuAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除菜单
func DelMenu(c *gin.Context) {
	var menu service.MenuId
	if err := c.ShouldBindJSON(&menu); err == nil {
		resCode := menu.DelMenu()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取菜单页
func GetMenu(c *gin.Context) {
	var menu service.MenuId
	id, err := strconv.Atoi(c.Query("id"))

	if id != 0 || err != nil {
		menu.ID = id
		if info, errCode := menu.GetMenu(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			//顶级菜单
			if info.ParentId == 0 {
				util.JsonSuccessResponse(c, info)
			} else {
				type MenuEditInfo struct {
					ParentMenu interface{} `json:"parent_menu"`
					Info       interface{} `json:"menu_info"`
				}
				ParentMenu, _ := model.GetMenu(map[string]interface{}{"id": info.ParentId})
				util.JsonSuccessResponse(c, MenuEditInfo{
					ParentMenu: ParentMenu,
					Info:       info,
				})
			}
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//保存顶级菜单
func SaveTopMenu(c *gin.Context) {
	var menu service.TopMenu
	if err := c.ShouldBindJSON(&menu); err == nil {
		resCode := menu.SaveTopMenu()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//保存子菜单
func SaveSubMenu(c *gin.Context) {
	var menu service.SubMenu
	if err := c.ShouldBindJSON(&menu); err == nil {
		resCode := menu.SaveSubMenu()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取菜单树
func GetTreeMenus(c *gin.Context) {
	list := service.GetTreeMenus()
	fmt.Println(list)
	util.JsonSuccessResponse(c, list)
}
