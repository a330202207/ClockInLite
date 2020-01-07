package routes

import (
	"ClockInLite/config"
	"ClockInLite/middleware/cors"
	"ClockInLite/package/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

//初始化路由
func InitRouter(e *gin.Engine) {
	//session
	//e.Use(session.Session())

	//使用 Logger 中间件
	e.Use(gin.Logger())

	//使用 Recovery 中间件
	e.Use(gin.Recovery())

	//使用日志中间件
	//e.Use(logger.LoggerToFile())

	//权限中间件
	//e.Use(casbin.CheckLoginHandle(
	//	casbin.AllowPathPrefixSkipper(casbin.NotCheckPermissionUrl()...),
	//))

	//跨域
	e.Use(cors.Cors())

	//设置环境
	gin.SetMode(config.ServerSetting.RunMode)

	//404
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": error.NOROUTE,
			"msg":  error.GetMsg(error.NOROUTE),
		})
		return
	})

	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": error.NOROUTE,
			"msg":  error.GetMsg(error.NOROUTE),
		})
		return
	})

	//注册API路由
	RegisterApiRouter(e)

	//注册后台路由
	RegisterAdminRouter(e)
}
