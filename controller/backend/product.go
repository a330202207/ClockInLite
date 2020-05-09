package backend

import (
	"ClockInLite/config"
	"ClockInLite/model"
	"ClockInLite/package/error"
	"ClockInLite/service"
	"ClockInLite/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

//商品列表
func GetProductList(c *gin.Context) {
	data := map[string]interface{}{}
	status, _ := strconv.Atoi(c.DefaultQuery("status", "0"))
	name := c.DefaultQuery("name", "")

	startTime, _ := strconv.Atoi(c.DefaultQuery("startTime", "0"))
	endTime, _ := strconv.Atoi(c.DefaultQuery("endTime", "0"))

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.ServerSetting.PageSize)))
	Offset := (page - 1) * pageSize

	if status != 0 {
		data["status"] = status
	}

	if name != "" {
		data["name"] = name
	}

	if startTime != 0 && endTime != 0 {
		data["created_at >="] = startTime
		data["created_at <="] = endTime
	}

	query, args, _ := util.WhereBuild(data)
	Product, count, _ := model.GetProductList(pageSize, Offset, "order_by asc,created_at", query, args...)

	util.JsonSuccessPage(c, count, Product)
}

//获取商品
func GetProduct(c *gin.Context) {
	var product service.ProductId
	id, err := strconv.Atoi(c.Query("id"))
	if id != 0 || err != nil {
		product.ID = id
		if info, errCode := product.GetProduct(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			util.JsonSuccessResponse(c, info)
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//添加商品
func AddProduct(c *gin.Context) {
	var product service.ProductInfo
	if err := c.ShouldBindJSON(&product); err == nil {
		resCode := product.AddProduct()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除商品
func DelProduct(c *gin.Context) {
	var ProductId service.ProductId
	if err := c.ShouldBindJSON(&ProductId); err == nil {
		resCode := ProductId.ProductDel()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//商品保存
func SaveProduct(c *gin.Context) {
	var product service.Product
	if err := c.ShouldBindJSON(&product); err == nil {
		resCode := product.ProductSave()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//商品上下架
func UpdateProductStatus(c *gin.Context) {
	var productStatus service.ProductStatus
	if err := c.ShouldBindJSON(&productStatus); err == nil {
		resCode := productStatus.UpdateProductStatus()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}
