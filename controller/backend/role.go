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

//角色列表页
func GetRoleList(c *gin.Context) {
	data := map[string]interface{}{"status <": 2}
	name := c.DefaultQuery("keyword", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.ServerSetting.PageSize)))
	Offset := (page - 1) * pageSize

	if name != "" {
		data["name like"] = "%" + name
	}

	query, args, _ := util.WhereBuild(data)
	roles, count, _ := model.GetRoleList(pageSize, Offset, "created_at desc", query, args...)

	util.JsonSuccessPage(c, count, roles)
}

//添加角色
func RoleAdd(c *gin.Context) {
	var role service.RoleMenu
	if err := c.BindJSON(&role); err == nil {
		resCode := role.RoleAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除角色
func RoleDel(c *gin.Context) {
	var role service.RoleId
	if err := c.BindJSON(&role); err == nil {
		resCode := role.RoleDel()
		fmt.Println(resCode)
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//编辑角色页
func RoleEdit(c *gin.Context) {
	var role service.RoleId
	id, err := strconv.Atoi(c.Query("id"))

	if id != 0 || err != nil {
		role.ID = id
		if info, errCode := role.RoleEdit(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			menus, _ := model.GetMenus(map[string]interface{}{})
			myMenus, _ := model.GetRoleMenus(map[string]interface{}{"role_id": id})

			type RoleEditInfo struct {
				Menus interface{} `json:"menus"`
				MyMenus   interface{} `json:"my_menus"`
				Info      interface{} `json:"role_info"`
			}

			util.JsonSuccessResponse(c, RoleEditInfo{
				Menus:   menus,
				MyMenus: myMenus,
				Info:    info,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

func MyMenus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	myMenus, _ := model.GetRoleMenus(map[string]interface{}{"role_id": id})
	util.JsonSuccessResponse(c, myMenus)
}

//保存角色
func RoleSave(c *gin.Context) {
	var role service.RoleInfo
	if err := c.BindJSON(&role); err == nil {
		resCode := role.RoleSave()
		util.HtmlResponse(c, resCode)
	} else {
		fmt.Println(err)
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}