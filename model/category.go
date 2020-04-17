package model

type Category struct {
	ID        int    `gorm:"primary_key" json:"id"`
	ParentId  int    `gorm:"default:0"   json:"parent_id"` //父级ID
	Name      string `gorm:"default:''"  json:"name"`      //分类名称
	OrderBy   int    `gorm:"default:'0'" json:"order_by"`  //菜单访问路由
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

//添加分类
func AddCategory(category *Category) (err error) {
	err = DB.Create(category).Error
	return
}

//获取分类列表
func GetCategoryList(Limit, Offset int, order string, query interface{}, args ...interface{}) (list []Category, count int, err error) {
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Find(&list).Error
	DB.Unscoped().Model(&Category{}).Where(query, args...).Count(&count)
	return
}

//菜单是否存在
func ExistCategory(maps interface{}) bool {
	var category Category
	DB.Unscoped().Where(maps).First(&category)
	if category.ID > 0 {
		return true
	}
	return false
}

//获取分类
func GetCategory(maps interface{}) (category Category, err error) {
	err = DB.Unscoped().Where(maps).First(&category).Error
	return
}

//获取多个分类
func GetCategories(maps interface{}) (list []Category, err error) {
	err = DB.Unscoped().Where(maps).Find(&list).Error
	return
}

//删除分类
func DelCategory(maps interface{}) (err error) {
	err = DB.Where(maps).Unscoped().Delete(&Category{}).Error
	return
}

//保存分类
func SaveCategory(id int, category Category) (err error) {
	err = DB.Model(&category).Where("id = ?", id).Updates(category).Error
	return
}
