package routes

import (
	"ClockInLite/controller/backend"
	"ClockInLite/package/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

//后台
func RegisterAdminRouter(e *gin.Engine) {

	e.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	admin := e.Group("/admin")
	{
		//登录
		admin.POST("/login", backend.Login)

		//登出
		admin.GET("/logout", backend.LogOut)

		//首页
		admin.GET("/index", backend.Index)

		//管理员列表
		admin.GET("/get/adminList", backend.GetAdminList)

		//添加管理员
		admin.POST("/add/admin", backend.AddAdmin)

		//删除管理员
		admin.POST("del/admin", backend.DelAdmin)

		//获取管理员
		admin.GET("/get/admin", backend.GetAdmin)

		//保存管理员信息
		admin.POST("/save/admin", backend.SaveAdmin)

		//角色列表
		admin.GET("/get/roleList", backend.GetRoleList)

		//添加角色
		admin.POST("/add/role", backend.AddRole)

		//删除角色
		admin.POST("/del/role", backend.DelRole)

		//获取所有角色
		admin.GET("/get/all/role", backend.GetAllRole)

		//保存角色
		admin.POST("/save/role", backend.SaveRole)

		//获取当前角色菜单
		admin.GET("/get/role/menus", backend.GetRoleMenus)

		//菜单列表
		admin.GET("/get/menuList", backend.GetMenuList)

		//获取菜单树结构
		admin.GET("/menu/menus", backend.GetTreeMenus)

		//添加顶级菜单
		admin.POST("/menu/add/topMenu", backend.TopMenuAdd)

		//添加子菜单
		admin.POST("/menu/add/subMenu", backend.SubMenuAdd)

		//删除菜单
		admin.POST("/del/menu", backend.DelMenu)

		//获取菜单页
		admin.GET("/get/menu", backend.GetMenu)

		//保存顶级菜单
		admin.POST("/save/topMenu", backend.SaveTopMenu)

		//保存子菜单
		admin.POST("/save/subMenu/", backend.SaveSubMenu)

		//添加分类
		admin.POST("/add/category", backend.AddCategory)

		//删除分类
		admin.POST("/del/category/", backend.DelCategory)

		//获取分类
		admin.GET("/get/category", backend.GetCategory)

		//保存分类
		admin.POST("/save/category", backend.SaveCategory)

		//获取分类列表
		admin.GET("/get/categoryList", backend.GetCategoryList)

		//获取多个分类
		admin.GET("/get/categories", backend.GetCategories)

		//获取商品列表
		admin.GET("/get/productList", backend.GetProductList)

		//获取商品
		admin.GET("/get/product", backend.GetProduct)

		//删除商品
		admin.POST("/del/product", backend.DelProduct)

		//保存商品
		admin.POST("/save/product", backend.SaveProduct)

		//商品上下架
		admin.POST("/updateStatus/product", backend.UpdateProductStatus)

		//添加商品
		admin.POST("/add/product", backend.AddProduct)

		//上传图片
		admin.POST("/upload/img", backend.UploadImg)

		//删除图片
		admin.POST("/del/img", backend.DelImg)

		//删除图片（软删除）
		admin.POST("/move/img", backend.MoveImg)
	}
}
