package casbin

import (
	"ClockInLite/model"
	"ClockInLite/util/convert"
	"fmt"
	"github.com/casbin/casbin"
)

var Enforcer *casbin.Enforcer

//初始化 角色-URL
func InitCasbin() (err error) {
	var enforcer *casbin.Enforcer
	casbinModel := `[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) == true \
			&& keyMatch2(r.obj, p.obj) == true \
			&& regexMatch(r.act, p.act) == true \
			|| r.sub == "root"`
	enforcer, err = casbin.NewEnforcerSafe(
		casbin.NewModel(casbinModel),
	)
	if err != nil {
		return
	}

	list, err := model.GetRoleMenus(map[string]interface{}{})
	if err != nil {
		return
	}

	for _, k := range list {
		setRolePermission(enforcer, k.RoleID, k.MenuID)
	}

	if len(list) == 0 {
		Enforcer = enforcer
		return
	}

	Enforcer = enforcer

	return
}

//设置角色权限
func setRolePermission(e *casbin.Enforcer, roleId, menuId int) {
	menu, err := model.GetMenu(map[string]interface{}{"id": menuId})
	if err != nil {
		return
	}
	fmt.Println(menu.MenuRouter)
	e.AddPermissionForUser(convert.ToString(roleId), menu.MenuRouter, "GET|POST")
}

//检查用户是否有权限
func CheckPermission(adminId, url, methodtype string) (bool, error) {
	return Enforcer.EnforceSafe(adminId, url, methodtype)
}

//用户角色处理
func AddRoleForUser(adminId int) (err error) {
	if Enforcer == nil {
		return
	}
	Enforcer.DeleteRolesForUser(convert.ToString(adminId))

	list, err := model.GetAdminRoles(map[string]interface{}{"admin_id": adminId})
	if err != nil {
		return
	}
	for _, v := range list {
		Enforcer.AddRoleForUser(convert.ToString(adminId), convert.ToString(v.RoleID))
	}
	return
}
