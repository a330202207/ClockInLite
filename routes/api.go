package routes

import (
	api "ClockInLite/controller/api/v1"
	"ClockInLite/middleware/jwt"
	"github.com/gin-gonic/gin"
)

//注册API模块路由
func RegisterApiRouter(e *gin.Engine) {
	apiRouter := e.Group("/api/v1")
	{
		apiRouter.GET("/test/index", api.Test)

		apiRouter.POST("test/getToken", api.GetToken)

		//apiRouter.POST("test/testToken", api.TestToken)
	}

	apiRouter.Use(jwt.JWT())
	{
		apiRouter.GET("test/testToken", api.TestToken)
	}
}
