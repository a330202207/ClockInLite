package model

type AdminRole struct {
	AdminID int `gorm:"default:0" json:"admin_id"`
	RoleID  int `gorm:"default:0" json:"role_id"`
}

//添加管理员-角色
func AddAdminRole(adminRole *AdminRole) (err error) {
	err = DB.Create(&adminRole).Error
	return
}

//获取管理员角色
func GetAdminRole(maps interface{}) (role AdminRole, err error) {
	err = DB.Unscoped().Where(maps).First(&role).Error
	return
}

//获取全部管理员-角色
func GetAdminRoles(maps interface{}) (roles AdminRole, err error) {
	err = DB.Unscoped().Where(maps).First(&roles).Error
	return
}

//删除管理员-角色
func DelAdminRole(adminID int) (err error) {
	err = DB.Model(&AdminRole{}).Where("admin_id = ?", adminID).Unscoped().Delete(&AdminRole{}).Error
	return
}

func SaveAdminRole(adminID int, roleID int) (err error) {
	err = DB.Model(&AdminRole{}).Where("admin_id = ?", adminID).Unscoped().Delete(&AdminRole{}).Error
	return
}
