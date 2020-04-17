package service

import (
	"ClockInLite/model"
	"ClockInLite/package/error"
)

type ProductId struct {
	ID int `form:"id" json:"id" binding:"required"`
}

type ProductInfo struct {
	Name       string   `form:"name"        json:"name"        binding:"required"`
	Details    string   `form:"details"     json:"details"     binding:"required"`
	Price      float64  `form:"price"       json:"price"       binding:"required"`
	CategoryId int      `form:"category_id" json:"category_id" binding:"required"`
	Num        int      `form:"num"         json:"num"         binding:"required"`
	Imgs       []string `form:"imgs"        json:"imgs"        binding:"required"`
	Status     int      `form:"status"      json:"status"`
	OrderBy    int      `form:"order_by"    json:"order_by"`
}

type Product struct {
	ProductId
	ProductInfo
}

//添加商品
func (productInfo *ProductInfo) AddProduct() int {
	whereMaps := map[string]interface{}{"name": productInfo.Name, "category_id": productInfo.CategoryId}
	isExist := model.ExistProduct(whereMaps)
	if isExist == true {
		return error.ERROR_EXIST_PRODUCT
	}
	product := model.Product{
		CategoryId: productInfo.CategoryId,
		Name:       productInfo.Name,
		Details:    productInfo.Details,
		Price:      productInfo.Price,
		Num:        productInfo.Num,
		Status:     productInfo.Status,
		OrderBy:    productInfo.OrderBy,
	}

	if err := model.AddProduct(&product, productInfo.Imgs); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//删除商品
func (productId *ProductId) ProductDel() int {
	id := map[string]interface{}{"id": productId.ID}
	isExist := model.ExistProduct(id)
	if isExist == false {
		return error.ERROR_NOT_EXIST_CATEGORY
	}

	err := model.DelProduct(productId.ID)
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}
	return error.SUCCESS
}

//获取商品
func (productId *ProductId) GetProduct() (model.Product, int) {
	productInfo, err := model.GetProduct(map[string]interface{}{"id": productId.ID})
	if err != nil {
		return productInfo, error.ERROR_NOT_EXIST_PRODUCT
	}
	return productInfo, error.SUCCESS
}

//保存商品
func (product *Product) ProductSave() int {
	id := product.ID
	productInfo := model.Product{
		CategoryId: product.CategoryId,
		Name:       product.Name,
		Details:    product.Details,
		Price:      product.Price,
		Num:        product.Num,
		Status:     product.Status,
		OrderBy:    product.OrderBy,
	}
	if err := model.SaveProduct(id, productInfo); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}
