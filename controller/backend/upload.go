package backend

import (
	"ClockInLite/package/error"
	"ClockInLite/package/upload"
	"ClockInLite/service"
	"ClockInLite/util"
	"github.com/gin-gonic/gin"
)

//图片上传
func UploadImg(c *gin.Context) {
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}

	if image != nil {
		imageName := upload.GetImageName(image.Filename)
		if !upload.CheckImageExt(imageName) {
			util.JsonErrResponse(c, error.ERROR_UPLOAD_IMAGE_EXT_ERR)
			return
		}

		if !upload.CheckImageSize(file) {
			util.JsonErrResponse(c, error.ERROR_UPLOAD_IMAGE_SIZE_ERR)
			return
		}

		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + imageName

		if err := c.SaveUploadedFile(image, src); err != nil {
			util.JsonErrResponse(c, error.ERROR_UPLOAD_IMAGE_FAIL)
			return
		}

		data := map[string]string{
			"imgUrl":  upload.GetImageFullUrl(imageName),
			"imgPath": savePath + imageName,
		}
		util.JsonSuccessResponse(c, data)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//图片删除
func DelImg(c *gin.Context) {
	var file service.File
	if err := c.ShouldBindJSON(&file); err == nil {
		resCode := file.DelImg()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除图片（软删除）
func MoveImg(c *gin.Context) {
	var file service.File
	if err := c.ShouldBindJSON(&file); err == nil {
		resCode := file.MoveImg()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}
