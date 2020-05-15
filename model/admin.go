package model

import (
	"time"
)

type Admin struct {
	Model

	UserName  string    `gorm:"default:''" json:"user_name"`    //管理员名称
	Password  string    `gorm:"default:''" json:"password"`     //密码
	Phone     string    `gorm:"default:''" json:"phone"`        //电话
	CreateIp  string    `gorm:"default:''" json:"create_ip"`    //创建时IP
	LoginIp   string    `gorm:"default:''" json:"login_ip"`     //登录时IP
	LoginDate time.Time `gorm:"default:null" json:"login_date"` //登录日期
	LoginCnt  int       `gorm:"default:0" json:"login_cnt"`     //登录次数
	Status    int       `gorm:"default:1" json:"status"`
}

//获取管理员
func GetAdmin(maps interface{}) (admin Admin, err error) {
	err = DB.Unscoped().Where(maps).First(&admin).Error
	return
}

//查询管理员是否存在
func ExistAdmin(maps interface{}) bool {
	var admin Admin
	DB.Unscoped().Where(maps).First(&admin)
	if admin.ID > 0 {
		return true
	}
	return false
}

//获取管理员列表
func GetAdminList(Limit, Offset int, order string, query interface{}, args ...interface{}) (list []Admin, count int, err error) {
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Find(&list).Error
	DB.Unscoped().Model(&Admin{}).Where(query, args...).Count(&count)
	return
}

//添加管理员
func AddAdmin(admin *Admin, roleId int) (err error) {

	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := DB.Create(admin).Error; err != nil {
		tx.Rollback()
		return err
	}

	adminRole := AdminRole{
		AdminID: admin.ID,
		RoleID:  roleId,
	}

	if err := AddAdminRole(&adminRole); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//删除管理员
func DelAdmin(id int) (err error) {

	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//删除管理员
	if err := DB.Where("id = ?", id).Unscoped().Delete(&Admin{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除关联角色
	if err := DB.Model(&AdminRole{}).Where("admin_id = ?", id).Unscoped().Delete(&AdminRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return
}

//保存管理员
func SaveAdmin(adminId int, roleID int, admin Admin) (err error) {
	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//保存管理员
	if err := DB.Unscoped().Model(&admin).Where("id = ?", adminId).Updates(admin).Error; err != nil {
		tx.Rollback()
		return err
	}

	//更新角色
	if err := SaveAdminRole(adminId, roleID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//更新登陆信息
func UpdateLoginInfo(id int, admin Admin) (err error) {
	err = DB.Unscoped().Model(&admin).Where("id = ?", id).Updates(admin).Error
	return
}
