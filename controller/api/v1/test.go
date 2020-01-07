package v1

import (
	"ClockInLite/config"
	"ClockInLite/package/error"
	"ClockInLite/util"
	"ClockInLite/util/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"mes":  "ok",
		"data": "",
	})
}

func GetToken(c *gin.Context) {
	Username := c.PostForm("username")
	password := c.PostForm("password")

	if Username == "" || password == "" {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
		return
	}

	timeOut := config.ServerSetting.JwtAdminTimeout
	token, expireTime, err := jwt.GenerateToken(Username, password, timeOut)
	if err != nil {
		util.JsonErrResponse(c, error.ERROR_AUTH_TOKEN)
		return
	}

	type info struct {
		Token string `json:"token"`
		ExpireTime int64 `json:"expire_Time"`
	}


	util.JsonSuccessResponse(c, info{
		Token:token,
		ExpireTime:expireTime,
	})
}

func TestToken(c *gin.Context) {
	token := c.Query("token")

	//body,_ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println("---body/--- \r\n "+string(body))
	fmt.Print(token)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"mes":  "token ok",
	})
}
