package routes

import (
	"ClockInLite/controller/backend"
	"github.com/gin-gonic/gin"
)

//后台
func RegisterAdminRouter(e *gin.Engine) {
	admin := e.Group("/admin")
	{
		//登录
		admin.POST("/login", backend.Login)

		//登出
		admin.GET("/logout", backend.LogOut)

		//首页
		admin.GET("/index", backend.Index)

		//管理员列表
		admin.GET("/admin/list", backend.GetAdminList)

		//添加管理员
		admin.POST("/admin/add", backend.AdminAdd)

		//删除管理员
		admin.POST("/admin/del", backend.AdminDel)

		//编辑管理员
		admin.GET("/admin/edit", backend.AdminEdit)

		//保存管理员信息
		admin.POST("/admin/save", backend.AdminSave)

		//角色列表
		admin.GET("/role/list", backend.GetRoleList)

		//添加角色
		admin.POST("/role/add", backend.RoleAdd)

		//删除角色
		admin.POST("/role/del", backend.RoleDel)

		//编辑角色页
		admin.GET("/role/edit", backend.RoleEdit)

		//保存角色
		admin.POST("/role/save", backend.RoleSave)

		//获取当前角色菜单
		admin.GET("/role/myMenus", backend.MyMenus)

		//菜单列表
		admin.GET("/menu/list", backend.GetMenuList)

		//获取菜单树结构
		admin.GET("/menu/menus", backend.GetTreeMenus)

		//添加顶级菜单
		admin.POST("/menu/addTop", backend.TopMenuAdd)

		//添加子菜单
		admin.POST("/menu/addSub", backend.SubMenuAdd)

		//删除菜单
		admin.POST("/menu/del", backend.MenuDel)

		//编辑菜单页
		admin.GET("/menu/edit", backend.MenuEdit)

		//保存顶级菜单
		admin.POST("/menu/saveTop", backend.TopMenuSave)

		//保存子菜单
		admin.POST("/menu/saveSub", backend.SubMenuSave)

		//添加分类
		admin.POST("/category/add", backend.AddCategory)

		//删除分类
		admin.POST("/category/del", backend.DelCategory)

		//编辑分类
		admin.GET("/category/edit", backend.EditCategory)

		//保存分类
		admin.POST("/category/save", backend.SaveCategory)

		//分类列表
		admin.GET("/category/list", backend.CategoryList)

	}
}
