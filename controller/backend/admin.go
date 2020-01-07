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

//管理员列表
func GetAdminList(c *gin.Context) {
	data := map[string]interface{}{"status <": 3}
	name := c.DefaultQuery("keyword", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.ServerSetting.PageSize)))
	Offset := (page - 1) * pageSize

	if name != "" {
		data["user_name like"] = name + "%"
	}

	query, args, _ := util.WhereBuild(data)
	admins, count, _ := model.GetAdminList(pageSize, Offset, "created_at desc", query, args...)

	util.JsonSuccessPage(c, count, admins)
}

//添加管理员
func AdminAdd(c *gin.Context) {
	var admin service.Admin
	admin.AdminInfo.CreateIp = c.ClientIP()
	if err := c.BindJSON(&admin); err == nil {
		resCode := admin.AdminAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除管理员
func AdminDel(c *gin.Context) {
	var AdminId service.AdminId
	if err := c.BindJSON(&AdminId); err == nil {
		resCode := AdminId.AdminDel()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}



//编辑管理员页
func AdminEdit(c *gin.Context) {
	var admin service.AdminId
	id, err := strconv.Atoi(c.Query("id"))

	if id != 0 || err != nil {
		admin.ID = id

		if info, errCode := admin.AdminEdit(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			//全部角色
			roles, _ := model.GetAllRoles()

			//我的角色
			myRoles, _ := model.GetAdminRoles(map[string]interface{}{"admin_id": admin.ID})
			type AdminEditInfo struct {
				RolesInfo interface{} `json:"roles"`
				MyRoles   interface{} `json:"my_roles"`
				Info      interface{} `json:"admin_info"`
			}

			util.JsonSuccessResponse(c, AdminEditInfo{
				RolesInfo:   roles,
				MyRoles: myRoles,
				Info:    info,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}

}

//保存管理员
func AdminSave(c *gin.Context) {
	var admin service.Account
	if err := c.BindJSON(&admin); err == nil {
		resCode := admin.AdminSave()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}
