package service

import (
	"ClockInLite/model"
	"ClockInLite/package/error"
)

type CategoryId struct {
	ID int `form:"id" json:"id" binding:"required"`
}

type CategoryInfo struct {
	ParentId int    `form:"parent_id" json:"parent_id"`
	Name     string `form:"name" json:"name" binding:"required"`
	OrderBy  int    `form:"order_by" json:"order_by"`
}

type Category struct {
	CategoryId
	CategoryInfo
}

//添加分类
func (categoryInfo *CategoryInfo) AddCategory() int {
	whereMaps := map[string]interface{}{"name": categoryInfo.Name, "parent_id": categoryInfo.ParentId}
	isExist := model.ExistCategory(whereMaps)
	if isExist == true {
		return error.ERROR_EXIST_CATEGORY
	}
	category := model.Category{
		ParentId: categoryInfo.ParentId,
		Name:     categoryInfo.Name,
		OrderBy:  categoryInfo.OrderBy,
	}
	if err := model.AddCategory(&category); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//删除分类
func (categoryId *CategoryId) CategoryDel() int {
	id := map[string]interface{}{"id": categoryId.ID}
	isExist := model.ExistCategory(id)
	if isExist == false {
		return error.ERROR_NOT_EXIST_CATEGORY
	}

	err := model.DelCategory(map[string]interface{}{"id": categoryId.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}
	return error.SUCCESS
}

//编辑分类
func (categoryId *CategoryId) CategoryEdit() (model.Category, int) {
	category, err := model.GetCategory(map[string]interface{}{"id": categoryId.ID})
	if err != nil {
		return category, error.ERROR_NOT_EXIST_CATEGORY
	}
	return category, error.SUCCESS
}

//保存分类
func (category *Category) CategorySave() int {
	id := category.ID
	categoryInfo := model.Category{
		ParentId: category.ParentId,
		Name:     category.Name,
		OrderBy:  category.OrderBy,
	}
	if err := model.SaveCategory(id, categoryInfo); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}
