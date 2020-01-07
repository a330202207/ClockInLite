package jwt

import (
	"ClockInLite/package/error"
	"ClockInLite/util"
	"ClockInLite/util/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		fmt.Println("token:", token)

		if token == "" {
			util.JsonErrResponse(c, error.INVALID_PARAMS)
			c.Abort()
			return
		} else {
			claims, err := jwt.ParseToken(token)
			fmt.Println("err:", err)
			if err != nil {
				util.JsonErrResponse(c, error.ERROR_AUTH_CHECK_TOKEN_FAIL)
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				util.JsonErrResponse(c, error.ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
