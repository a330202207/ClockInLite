package model

type Role struct {
	Model

	Name   string `gorm:"default:''" json:"name"` //角色名称
	Status int    `gorm:"default:1" json:"status"`
}

//获取角色
func GetRole(maps interface{}) (role Role, err error) {
	err = DB.Unscoped().Where(maps).First(&role).Error
	return
}

//角色是否存在
func ExistRole(maps interface{}) bool {
	var role Role
	DB.Unscoped().Where(maps).First(&role)
	if role.ID > 0 {
		return true
	}
	return false
}

//角色列表
func GetRoleList(Limit, Offset int, order string, query interface{}, args ...interface{}) (list []Role, count int, err error) {
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Find(&list).Error
	DB.Unscoped().Model(&Role{}).Where(query, args...).Count(&count)
	return
}

//获取全部角色
func GetAllRoles() (roles []Role, err error) {
	err = DB.Unscoped().Where("status = ?", 1).Find(&roles).Error
	return
}

//添加角色
func AddRole(role *Role) (id int, err error) {
	err = DB.Create(role).Error
	id = role.ID
	return
}

//删除角色
func DelRole(maps interface{}) (err error) {
	err = DB.Model(&Role{}).Unscoped().Where(maps).Update("status", 2).Error
	return
}

//保存角色
func SaveRole(id int, role Role) (err error) {
	err = DB.Model(&role).Where("id = ?", id).Updates(role).Error
	return
}

//修改角色状态
func UpdateRoleStatus(id int, role Role) (err error) {
	err = DB.Model(&role).Where("id = ?", id).Unscoped().Updates(role).Error
	return
}
