package model

type Menu struct {
	ID         int    `gorm:"primary_key" json:"id"`
	ParentId   int    `gorm:"default:0"   json:"parent_id"`   //上级ID
	Name       string `gorm:"default:''"  json:"name"`        //菜单名称
	MenuRouter string `gorm:"default:''"  json:"menu_router"` //菜单访问路由
	OrderBy    int    `gorm:"default:'0'" json:"order_by"`    //排序
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
}

type TreeMenus struct {
	ID         int          `json:"key"`
	ParentId   int          `json:"parent_id"`   //上级ID
	Name       string       `json:"title"`       //菜单名称
	MenuRouter string       `json:"menu_router"` //菜单访问路由
	OrderBy    int          `json:"order_by"`    //菜单访问路由
	Children   []*TreeMenus `json:"children"`
}

//添加菜单
func AddMenu(menu *Menu) (err error) {
	err = DB.Create(menu).Error
	return
}

//获取菜单列表
func GetMenuList(Limit, Offset int, order string, query interface{}, args ...interface{}) (list []Menu, count int, err error) {
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Find(&list).Error
	DB.Unscoped().Where(query, args...).Model(&Menu{}).Count(&count)
	return
}

//获取菜单
func GetMenu(maps interface{}) (menu Menu, err error) {
	err = DB.Unscoped().Where(maps).First(&menu).Error
	return
}

//获取多个菜单
func GetMenus(maps interface{}) (menu []Menu, err error) {
	err = DB.Unscoped().Where(maps).Order("order_by asc").Find(&menu).Error
	return
}

//获取菜单树
func (menu *Menu) GetTreeMenus(parentId int) []*TreeMenus {
	menus, _ := GetMenus(map[string]interface{}{"parent_id": parentId})

	treeList := []*TreeMenus{}

	for _, v := range menus {

		child := v.GetTreeMenus(v.ID)
		node := &TreeMenus{
			ID:         v.ID,
			ParentId:   v.ParentId,
			Name:       v.Name,
			MenuRouter: v.MenuRouter,
			OrderBy:    v.OrderBy,
		}
		node.Children = child

		treeList = append(treeList, node)
	}

	return treeList
}

//菜单是否存在
func ExistMenu(maps interface{}) bool {
	var menu Menu
	DB.Unscoped().Where(maps).First(&menu)
	if menu.ID > 0 {
		return true
	}
	return false
}

//删除菜单
func DelMenu(maps interface{}) (err error) {
	err = DB.Where(maps).Unscoped().Delete(&Menu{}).Error
	return
}

//保存菜单
func SaveMenu(id int, menu Menu) (err error) {
	err = DB.Model(&menu).Where("id = ?", id).Updates(menu).Error
	return
}

//获取管理员菜单
func GetMenuByAdminId(adminId int) (menu []Menu, err error) {
	sql := `select * from api_menu
	      where id in (
					select menu_id from api_role_menu where 
				  role_id in (select role_id from api_admin_role where admin_id=?)
				)`
	err = DB.Raw(sql, adminId).Find(&menu).Error
	return
}
