package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := cors.DefaultConfig()

		config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
		config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
		if gin.Mode() == gin.ReleaseMode {
			// 生产环境需要配置跨域域名，否则403
			config.AllowOrigins = []string{"http://www.example.com"}
		} else {
			config.AllowAllOrigins = true
		}
		config.AllowCredentials = true
		cors.New(config)
		c.Next()
	}
}
