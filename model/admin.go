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
func AddAdmin(admin *Admin) (id int, err error) {
	err = DB.Create(admin).Error
	id = admin.ID
	return
}

//删除管理员
func DelAdmin(maps interface{}) (err error) {
	err = DB.Model(&Admin{}).Unscoped().Where(maps).Update("status", 3).Error
	return
}

//保存管理员
func SaveAdmin(id int, admin Admin) (err error) {
	err = DB.Unscoped().Model(&admin).Where("id = ?", id).Updates(admin).Error
	return
}

//更新登陆信息
func UpdateLoginInfo(id int, admin Admin) (err error) {
	err = DB.Unscoped().Model(&admin).Where("id = ?", id).Updates(admin).Error
	return
}
