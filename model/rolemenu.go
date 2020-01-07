package model

import (
	"ClockInLite/util"
	"fmt"
	"strconv"
	"strings"
)

type RoleMenu struct {
	RoleID int `gorm:"default:0" json:"role_id"`
	MenuID int `gorm:"default:0" json:"menu_id"`
}

//添加角色-菜单
func AddRoleMenu(roleID int, menuIDs []string) (err error) {
	str := []string{"role_id", "menu_id"}
	newArr := [][]string{}
	for _, v := range menuIDs {
		Arr := append([]string{strconv.Itoa(roleID)}, v)
		newArr = append(newArr, Arr)
	}

	var newStr string
	for _, v := range newArr {
		newStr += fmt.Sprintf("('%s'),", strings.Join(v, "','"))
	}
	key := strings.Join(str, ",")
	val := strings.TrimRight(newStr, ",")

	sql := util.BatchInsert("api_role_menu", key, val)

	err = DB.Exec(sql).Error

	return
}

//获取全部角色-菜单
func GetRoleMenus(maps interface{}) (list []RoleMenu, err error) {
	err = DB.Unscoped().Where(maps).Find(&list).Error
	return
}

//删除角色-菜单
func DelRoleMenu(roleID int) (err error) {
	err = DB.Model(&AdminRole{}).Where("role_id = ?", roleID).Unscoped().Delete(&RoleMenu{}).Error
	return
}
