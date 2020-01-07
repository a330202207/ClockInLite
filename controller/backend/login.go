package backend

import (
	"ClockInLite/config"
	"ClockInLite/middleware/casbin"
	"ClockInLite/model"
	"ClockInLite/package/error"
	"ClockInLite/service"
	"ClockInLite/util"
	"ClockInLite/util/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//登陆
func Login(c *gin.Context) {
	var accountInfo service.LoginInfo
	if err := c.Bind(&accountInfo); err != nil {
		util.JsonErrResponse(c, error.ERROR_NOT_EXIST_USER)
		return
	}

	//检查登陆
	admin, errCode := accountInfo.CheckLogin()
	if errCode != 200 {
		util.JsonErrResponse(c, errCode)
		return
	}

	nowLoginCnt := admin.LoginCnt
	loginInfo := model.Admin{
		LoginDate: time.Now(),
		LoginIp:   c.ClientIP(),
		LoginCnt:  nowLoginCnt + 1,
	}

	//更新登陆信息
	if err := model.UpdateLoginInfo(admin.ID, loginInfo); err != nil {
		util.JsonErrResponse(c, error.ERROR_SQL_UPDATE_FAIL)
	}

	//添加当前用户角色
	err := casbin.AddRoleForUser(admin.ID)
	if err != nil {
		util.JsonErrResponse(c, error.ERROR)
	}

	//获取Token
	token, _, _ := jwt.GenerateToken(
		accountInfo.UserName.UserName,
		accountInfo.Password.Password,
		config.ServerSetting.JwtAdminTimeout)

	loggedInfo := service.LoggedInfo{
		Token:     token,
		AdminID:   admin.ID,
		AdminName: admin.UserName,
		Menus:     service.GetAdminMenus(admin.ID),
	}

	util.JsonSuccessResponse(c, loggedInfo)
}

//登出
func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": error.SUCCESS,
		"msg":  error.GetMsg(error.SUCCESS),
	})
}
