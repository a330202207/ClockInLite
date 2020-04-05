package model

import (
	"ClockInLite/util"
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	ID         int     `gorm:"primary_key" json:"id"`
	CategoryId int     `gorm:"default:0"   json:"category_id"` //商品分类ID
	Name       string  `gorm:"default:''"  json:"name"`        //商品名称
	OrderBy    int     `gorm:"default:'0'" json:"order_by"`    //排序
	Price      float64 `gorm:"default:'0'" json:"price"`       //价格
	Num        int     `gorm:"default:'0'" json:"num"`         //库存
	Status     int     `gorm:"default:'0'" json:"status"`      //状态
	Details    string  `gorm:"default:'0'" json:"details"`     //详情
	CreatedAt  int     `json:"created_at"`
	UpdatedAt  int     `json:"updated_at"`
}

//添加商品
func AddProduct(product *Product, imgs []string) (err error) {
	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := DB.Create(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	ProductId := product.ID

	if err := addProductImg(ProductId, imgs); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func addProductImg(productId int, imgs []string) (err error) {
	str := []string{"admin_id", "role_id"}
	newArr := [][]string{}
	for _, v := range imgs {
		Arr := append([]string{strconv.Itoa(productId)}, v)
		newArr = append(newArr, Arr)
	}

	var newStr string
	for _, v := range newArr {

		newStr += fmt.Sprintf("('%s'),", strings.Join(v, "','"))

	}
	key := strings.Join(str, ",")
	val := strings.TrimRight(newStr, ",")

	sql := util.BatchInsert("api_product_img", key, val)
	fmt.Println("sql", sql)
	//err = DB.Exec(sql).Error

	return
}

//删除商品
func DelProduct(id int) (err error) {
	err = DB.Where("id = ?").Unscoped().Delete(&Product{}).Error
	return
}

//商品是否存在
func ExistProduct(maps interface{}) bool {
	var product Product
	DB.Unscoped().Where(maps).First(&product)
	if product.ID > 0 {
		return true
	}
	return false
}

//商品列表
func GetProductList(Limit, Offset int, order string, query interface{}, args ...interface{}) (list []Product, count int, err error) {
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Find(&list).Error
	DB.Unscoped().Model(&Product{}).Where(query, args...).Count(&count)
	return
}

//保存商品
func SaveProduct(id int, product Product) (err error) {
	err = DB.Model(&product).Where("id = ?", id).Updates(product).Error
	return
}
