package model

import (
	"ClockInLite/util"
	"fmt"
	"strconv"
	"strings"
)

type AdminRole struct {
	AdminID int `gorm:"default:0" json:"admin_id"`
	RoleID  int `gorm:"default:0" json:"role_id"`
}

//添加管理员-角色
func AddAdminRole(adminId int, roleIds []string) (err error) {
	str := []string{"admin_id", "role_id"}
	newArr := [][]string{}
	for _, v := range roleIds {
		Arr := append([]string{strconv.Itoa(adminId)}, v)
		newArr = append(newArr, Arr)
	}

	var newStr string
	for _, v := range newArr {

		newStr += fmt.Sprintf("('%s'),", strings.Join(v, "','"))

	}
	key := strings.Join(str, ",")
	val := strings.TrimRight(newStr, ",")

	sql := util.BatchInsert("api_admin_role", key, val)

	err = DB.Exec(sql).Error

	return
}

func GetAdminRole(maps interface{}) (role AdminRole, err error) {
	err = DB.Unscoped().Where(maps).First(&role).Error
	return
}

//获取全部管理员-角色
func GetAdminRoles(maps interface{}) (roles []AdminRole, err error) {
	err = DB.Unscoped().Where(maps).Find(&roles).Error
	return
}

//删除管理员-角色
func DelAdminRole(adminID int) (err error) {
	err = DB.Model(&AdminRole{}).Where("admin_id = ?", adminID).Unscoped().Delete(&AdminRole{}).Error
	return
}
