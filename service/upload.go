package service

import (
	"ClockInLite/config"
	"ClockInLite/package/error"
	"ClockInLite/package/file"
	"ClockInLite/package/upload"
	"fmt"
	"strings"
)

type File struct {
	Url []string `form:"url" json:"urls" binding:"required"`
}

//删除图片
func (img *File) DelImg() int {
	urls := img.Url
	savePath := upload.GetImagePath()
	for _, val := range urls {
		path := config.AppSetting.StorageRootPath + val[strings.Index(val, savePath):]
		//fmt.Println("path:",path)
		if file.CheckExist(path) {
			break
			//return error.ERROR_NOT_EXIST_IMAGE
		}

		if err := file.DelFile(path); err != nil {
			break
			//return error.ERROR_DEL_IMAGE_FAIL
		}
	}

	return error.SUCCESS
}

//删除图片（软删除）
func (img *File) MoveImg() int {
	urls := img.Url
	savePath := upload.GetImagePath()
	fmt.Println("savePath:", savePath)
	for _, val := range urls {
		src := config.AppSetting.StorageRootPath + val[strings.Index(val, savePath):]

		if file.CheckExist(src) {
			break
			//return error.ERROR_NOT_EXIST_IMAGE
		}

		toPath := config.AppSetting.StorageRootPath + config.FileSetting.ImageMovePath + val[strings.Index(val, savePath):]
		if err := file.MoveFile(src, toPath); err != nil {
			break
			//return error.ERROR_DEL_IMAGE_FAIL
		}
	}
	return error.SUCCESS
}
