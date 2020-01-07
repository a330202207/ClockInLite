package backend

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CurrentLoginUserRole struct {
	HasPermission func(string) bool
}

//首页
func Index(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	c.HTML(http.StatusOK, "backend_index.html", gin.H{
		"user_id": userId,
	})

	//session := sessions.Default(c)
	//
	//currentUserRole := session.Get("user_id").(string)
	//c.HTML(http.StatusOK, "index.html", gin.H{
	//	"user":CurrentLoginUserRole{
	//		HasPermission: func(sys_res_id string) bool {
	//			return casbinObj.Enforce(currentUserRole,sys_res_id,"GET")
	//		},
	//	},
	//})

}
