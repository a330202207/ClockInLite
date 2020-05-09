package model

import (
	"ClockInLite/util"
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	Model

	CategoryId   int          `gorm:"default:0"   json:"category_id"`               //商品分类ID
	CategoryInfo Category     `gorm:"FOREIGNKEY:CategoryId"   json:"category_info"` //分类名称
	Name         string       `gorm:"default:''"  json:"name"`                      //商品名称
	OrderBy      int          `gorm:"default:'0'" json:"order_by"`                  //排序
	Price        float64      `gorm:"default:'0'" json:"price"`                     //价格
	Num          int          `gorm:"default:'0'" json:"num"`                       //库存
	Status       int          `gorm:"default:'0'" json:"status"`                    //状态
	Details      string       `gorm:"default:'0'" json:"details"`                   //详情
	Img          []ProductImg `gorm:"FOREIGNKEY:ProductId" json:"imgs"`             //One-To-Many (拥有多个 - ProductImg表的ProductID作外键)
}

type ProductImg struct {
	ID        int    `gorm:"primary_key" json:"id"`
	ProductId int    `gorm:"default:0"   json:"product_id"` //商品ID
	Img       string `gorm:"default:''"  json:"img"`        //商品图片
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
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

//添加商品图片
func addProductImg(productId int, imgs []string) (err error) {
	str := []string{"product_id", "img"}
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
	err = DB.Exec(sql).Error
	return
}

//删除商品
func DelProduct(id int) (err error) {
	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//删除商品
	if err := DB.Where("id = ?", id).Unscoped().Delete(&Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除图片
	if err := DB.Where("product_id = ?", id).Unscoped().Delete(&ProductImg{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Preload("Img").Preload("CategoryInfo").Find(&list).Error
	DB.Unscoped().Model(&Product{}).Where(query, args...).Count(&count)
	return
}

//获取商品信息
func GetProduct(maps interface{}) (product Product, err error) {
	DB.Unscoped().Where(maps).First(&product)
	DB.Model(&product).Related(&product.Img, "ProductId")
	err = DB.Model(&product).Related(&product.CategoryInfo, "CategoryId").Error
	return
}

//保存商品
func SaveProduct(id int, product Product, imgs []string) (err error) {
	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := DB.Model(&product).Where("id = ?", id).Unscoped().Updates(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除图片
	if err := DB.Where("product_id = ?", id).Unscoped().Delete(&ProductImg{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//添加图片
	if err := addProductImg(id, imgs); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//商品上下架
func UpdateProductStatus(id int, product Product) (err error) {
	err = DB.Model(&product).Where("id = ?", id).Unscoped().Updates(product).Error
	return
}
