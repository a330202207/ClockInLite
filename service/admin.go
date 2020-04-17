package service

import (
	"ClockInLite/model"
	"ClockInLite/package/error"
	"golang.org/x/crypto/bcrypt"
)

type LoginInfo struct {
	UserName
	Password
}

type Admin struct {
	LoginInfo
	AdminInfo
}

type Account struct {
	AdminId
	UserName
	AdminInfo
}

type AdminId struct {
	ID int `form:"id" json:"id" binding:"required"`
}

type UserName struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
}

type Password struct {
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

type AdminInfo struct {
	Phone    string   `json:"phone"`
	CreateIp string   `json:"login_ip"`
	RoleIDs  []string `form:"role_ids" json:"role_ids" binding:"required"`
	Status   int      `form:"status" json:"status" binding:"required"`
}

type LoggedInfo struct {
	Token     string       `json:"token"`
	AdminID   int          `json:"user_id"`
	AdminName string       `json:"user_name"`
	Menus     []model.Menu `json:"menus"`
}

//登陆
func (account *LoginInfo) CheckLogin() (model.Admin, int) {
	userName := map[string]interface{}{"user_name": account.UserName.UserName}
	admin, err := model.GetAdmin(userName)
	if err != nil {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//密码不对
	if CheckPassword(admin.Password, account.Password.Password) == false {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//禁用账户
	if CheckStatus(admin.Status) == false {
		return admin, error.ERROR_DISABLE_USER
	}

	return admin, error.SUCCESS
}

//检查用户密码
func CheckPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

//检查用户状态
func CheckStatus(status int) bool {
	if status == 2 || status == 3 {
		return false
	}
	return true
}

//添加用户
func (admin *Admin) AddAdmin() int {
	whereMap := map[string]interface{}{"user_name": admin.UserName.UserName, "status": 1}
	isExist := model.ExistAdmin(whereMap)
	if isExist == true {
		return error.ERROR_EXIST_USER
	}

	hashPassword, ok := SetPassword(admin.Password.Password)
	if ok == false {
		return error.ERROR_PASSWORD_USER
	}

	adminInfo := model.Admin{
		UserName: admin.UserName.UserName,
		CreateIp: admin.CreateIp,
		Password: hashPassword,
	}

	adminId, err := model.AddAdmin(&adminInfo)
	if err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	if model.AddAdminRole(adminId, admin.AdminInfo.RoleIDs) != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	return error.SUCCESS
}

//设置密码
func SetPassword(password string) (string, bool) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", false
	}
	return string(hash), true
}

//删除管理员
func (adminId *AdminId) DelAdmin() int {
	whereMap := map[string]interface{}{"id": adminId.ID, "status": 1}
	isExist := model.ExistAdmin(whereMap)
	if isExist == false {
		return error.ERROR_NOT_EXIST_USER
	}

	err := model.DelAdmin(map[string]interface{}{"id": adminId.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}

	if err := model.DelAdminRole(adminId.ID); err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}

	return error.SUCCESS
}

//获取管理员
func (adminId *AdminId) GetAdmin() (model.Admin, int) {
	adminInfo, err := model.GetAdmin(map[string]interface{}{"id": adminId.ID})
	if err != nil {
		return adminInfo, error.ERROR
	}
	return adminInfo, error.SUCCESS
}

//保存管理员
func (account *Account) SaveAdmin() int {
	id := account.ID
	admin := model.Admin{
		UserName: account.UserName.UserName,
		Status:   account.AdminInfo.Status,
	}
	if err := model.SaveAdmin(id, admin); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}

	if err := model.DelAdminRole(id); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}

	if err := model.AddAdminRole(id, account.AdminInfo.RoleIDs); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}

	return error.SUCCESS
}
