package backend

import (
	"ClockInLite/config"
	"ClockInLite/model"
	"ClockInLite/package/error"
	"ClockInLite/service"
	"ClockInLite/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

//角色列表页
func GetRoleList(c *gin.Context) {
	data := map[string]interface{}{}
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
func AddRole(c *gin.Context) {
	var role service.RoleMenu
	if err := c.ShouldBindJSON(&role); err == nil {
		resCode := role.AddRole()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除角色
func DelRole(c *gin.Context) {
	var role service.RoleId
	if err := c.ShouldBindJSON(&role); err == nil {
		resCode := role.DelRole()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取角色菜单
func GetRoleMenus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	myMenus, _ := model.GetRoleMenus(map[string]interface{}{"role_id": id})
	util.JsonSuccessResponse(c, myMenus)
}

//获取所有角色
func GetAllRole(c *gin.Context) {
	//全部角色
	roles, _ := model.GetAllRoles()
	util.JsonSuccessResponse(c, roles)
}

//保存角色
func SaveRole(c *gin.Context) {
	var role service.RoleInfo
	if err := c.ShouldBindJSON(&role); err == nil {
		resCode := role.SaveRole()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}
