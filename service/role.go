package service

import (
	"ClockInLite/model"
	"ClockInLite/package/error"
)

type RoleId struct {
	ID int `form:"id" json:"id" binding:"required"`
}

type RoleName struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type RoleInfo struct {
	RoleId
	RoleMenu
}

type RoleMenu struct {
	RoleName
	MenuIDs []string `form:"menu_ids" json:"menu_ids" binding:"required"`
}

//添加角色
func (roleInfo *RoleMenu) AddRole() int {
	name := map[string]interface{}{"name": roleInfo.Name}
	isExist := model.ExistRole(name)

	if isExist == true {
		return error.ERROR_EXIST_ROLE
	}
	role := model.Role{
		Name: roleInfo.Name,
	}

	roleID, err := model.AddRole(&role)
	if err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	if model.AddRoleMenu(roleID, roleInfo.MenuIDs) != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	return error.SUCCESS
}

//删除角色
func (role *RoleId) DelRole() int {
	whereMap := map[string]interface{}{"id": role.ID}
	isExist := model.ExistRole(whereMap)
	if isExist == false {
		return error.ERROR_SQL_DELETE_FAIL
	}

	err := model.DelRole(role.ID)
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}

	return error.SUCCESS
}

//获取角色
func (role *RoleId) GetRole() (model.Role, int) {
	roleInfo, err := model.GetRole(map[string]interface{}{"id": role.ID})
	if err != nil {
		return roleInfo, error.ERROR_NOT_EXIST_ROLE
	}
	return roleInfo, error.SUCCESS
}

//保存角色
func (roleInfo *RoleInfo) SaveRole() int {
	id := roleInfo.ID
	role := model.Role{
		Name: roleInfo.Name,
	}
	if err := model.SaveRole(id, role, roleInfo.MenuIDs); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}

	return error.SUCCESS
}
