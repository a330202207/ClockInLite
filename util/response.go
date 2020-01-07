package util

import (
	"ClockInLite/package/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Context map[string]interface{}

type ResponseModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseModelBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ResponsePageData struct {
	Total   int      `json:"total"`
	List    interface{} `json:"list"`
}

//json返回格式
func JsonResponse(c *gin.Context, httpCode int, v interface{}) {
	c.JSON(httpCode, v)
}

func JsonSuccessPage(c *gin.Context, total int, list interface{}) {
	ret := ResponseModel{
		Code: error.SUCCESS,
		Msg:  error.GetMsg(error.SUCCESS),
		Data: ResponsePageData{
			Total:   total,
			List:    list,
		},
	}

	JsonResponse(c, http.StatusOK, &ret)
}

// 响应成功
func JsonSuccessResponse(c *gin.Context, data interface{}) {

	ret := ResponseModel{
		Code: error.SUCCESS,
		Msg:  error.GetMsg(error.SUCCESS),
		Data: data,
	}
	JsonResponse(c, http.StatusOK, &ret)
}

// 响应失败
func JsonErrResponse(c *gin.Context, errCode int) {
	ret := ResponseModelBase{
		Code: errCode,
		Msg:  error.GetMsg(errCode),
	}
	JsonResponse(c, http.StatusOK, &ret)
}

func HtmlResponse(c *gin.Context, errCode int) {
	c.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg":  error.GetMsg(errCode),
	})
}
