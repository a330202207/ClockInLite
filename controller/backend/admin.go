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
func AddAdmin(c *gin.Context) {
	var admin service.Admin
	admin.AdminInfo.CreateIp = c.ClientIP()
	if err := c.ShouldBindJSON(&admin); err == nil {
		resCode := admin.AddAdmin()
		util.HtmlResponse(c, resCode)
	} else {
		fmt.Println(err)
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除管理员
func DelAdmin(c *gin.Context) {
	var AdminId service.AdminId
	if err := c.ShouldBindJSON(&AdminId); err == nil {
		resCode := AdminId.DelAdmin()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取管理员
func GetAdmin(c *gin.Context) {
	var admin service.AdminId
	id, err := strconv.Atoi(c.Query("id"))

	if id != 0 || err != nil {
		admin.ID = id

		if info, errCode := admin.GetAdmin(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {

			//我的角色
			myRoles, _ := model.GetAdminRole(map[string]interface{}{"admin_id": admin.ID})
			type AdminEditInfo struct {
				MyRoles interface{} `json:"my_role"`
				Info    interface{} `json:"admin_info"`
			}

			util.JsonSuccessResponse(c, AdminEditInfo{
				MyRoles: myRoles,
				Info:    info,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}

}

//保存管理员
func SaveAdmin(c *gin.Context) {
	var admin service.Account
	if err := c.ShouldBindJSON(&admin); err == nil {
		resCode := admin.SaveAdmin()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}
