package backend

import (
	"ClockInLite/config"
	"ClockInLite/model"
	"ClockInLite/package/error"
	"ClockInLite/service"
	"ClockInLite/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//菜单列表页
func GetMenuList(c *gin.Context) {
	data := map[string]interface{}{}
	parentId, _ := strconv.Atoi(c.DefaultQuery("parent_id", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.ServerSetting.PageSize)))

	Offset := (page - 1) * pageSize
	if parentId != 0 {
		data["parent_id"] = parentId
	} else {
		data["parent_id"] = 0
	}
	query, args, _ := util.WhereBuild(data)
	Menus, count, _ := model.GetMenuList(pageSize, Offset, "order_by asc,created_at", query, args)

	util.JsonSuccessPage(c, count, Menus)
}

//添加菜单
func AddMenu(c *gin.Context) {
	var menu service.MenuInfo
	if err := c.ShouldBindJSON(&menu); err == nil {
		resCode := menu.AddMenu()
		util.HtmlResponse(c, resCode)
	} else {
		fmt.Println(err)
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

//保存菜单
func SaveMenu(c *gin.Context) {
	var menu service.Menu
	if err := c.ShouldBindJSON(&menu); err == nil {
		resCode := menu.SaveMenu()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取菜单树
func GetTreeMenus(c *gin.Context) {
	var menu model.Menu
	list := menu.GetTreeMenus(0)

	type DataRes struct {
		Code int                `json:"code"`
		Msg  string             `json:"msg"`
		Data []*model.TreeMenus `json:"data"`
	}

	util.JsonResponse(c, http.StatusOK, DataRes{
		Code: error.SUCCESS,
		Msg:  error.GetMsg(error.SUCCESS),
		Data: list,
	})
}
