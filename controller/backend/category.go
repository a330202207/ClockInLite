package backend

import (
	"ClockInLite/config"
	"ClockInLite/model"
	"ClockInLite/package/error"
	"ClockInLite/service"
	"ClockInLite/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var category service.CategoryInfo
	if err := c.ShouldBindJSON(&category); err == nil {
		resCode := category.AddCategory()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除分类
func DelCategory(c *gin.Context) {
	var CategoryId service.CategoryId
	if err := c.ShouldBindJSON(&CategoryId); err == nil {
		resCode := CategoryId.DelCategory()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取分类
func GetCategory(c *gin.Context) {
	var CategoryId service.CategoryId
	id, err := strconv.Atoi(c.Query("id"))
	if id != 0 || err != nil {
		CategoryId.ID = id
		if info, errCode := CategoryId.GetCategory(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			util.JsonSuccessResponse(c, info)
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//保存分类
func SaveCategory(c *gin.Context) {
	var category service.Category
	if err := c.ShouldBindJSON(&category); err == nil {
		resCode := category.SaveCategory()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//分类列表
func GetCategoryList(c *gin.Context) {
	data := map[string]interface{}{}
	parentId, _ := strconv.Atoi(c.DefaultQuery("parent_id", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.ServerSetting.PageSize)))
	Offset := (page - 1) * pageSize

	if parentId != 0 {
		data["parent_id"] = parentId
	} else {
		data["parent_id"] = 0
	}

	query, args, _ := util.WhereBuild(data)
	Category, count, _ := model.GetCategoryList(pageSize, Offset, "id asc", query, args...)

	util.JsonSuccessPage(c, count, Category)
}

//获取多个分类
func GetCategories(c *gin.Context) {
	parentId, _ := strconv.Atoi(c.DefaultQuery("parent_id", "0"))
	list, _ := model.GetCategories(map[string]interface{}{"parent_id": parentId})

	util.JsonSuccessResponse(c, list)
}
