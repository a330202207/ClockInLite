package casbin

import (
	"ClockInLite/package/error"
	"ClockInLite/util"
	"ClockInLite/util/convert"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool // SkipperFunc 定义中间件跳过函数

func NotCheckPermissionUrl() (notCheckPermissionUrlArr []string) {

	apiPrefix := "/admin"

	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, "/static")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/backend_login.html")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/login.html")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/logout")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/backend_index.html")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/menu/menus")
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/role/myMenus")
	return
}

//检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		for _, p := range prefixes {
			if strings.HasPrefix(path, p) {
				return true
			}
		}
		return false
	}
}

func CheckLoginHandle(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		//获取请求的URL
		objUrl := c.Request.URL.RequestURI()

		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		//获取请求方法
		act := c.Request.Method

		//获取用户角色
		s := sessions.Default(c)

		adminId := s.Get("admin_id")

		if nil == adminId {
			c.Redirect(http.StatusTemporaryRedirect, "/admin/backend_login.html")
			return
		}
		res, err := CheckPermission(convert.ToString(adminId), objUrl, act)
		//判断权限
		if err != nil {
			util.JsonErrResponse(c, error.ERROR)
			c.Abort()
			return
		} else if !res {
			util.JsonErrResponse(c, error.NOPERMISSION)
			fmt.Println("err:", err, "objUrl:", objUrl, "act:", act)
			c.Abort()
			return
		}
		fmt.Println("res:", res)
		c.Next()
	}
}
