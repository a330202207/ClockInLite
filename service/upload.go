package service

import (
	"ClockInLite/config"
	"ClockInLite/package/error"
	"ClockInLite/package/file"
)

type File struct {
	Path string `form:"path" json:"path" binding:"required"`
}

//删除图片
func (img *File) DelImg() int {
	path := config.AppSetting.StorageRootPath + img.Path
	if file.CheckExist(path) {
		return error.ERROR_NOT_EXIST_IMAGE
	}

	if err := file.DelFile(path); err != nil {
		return error.ERROR_DEL_IMAGE_FAIL
	}
	return error.SUCCESS
}

//删除图片（软删除）
func (img *File) MoveImg() int {
	src := config.AppSetting.StorageRootPath + img.Path
	if file.CheckExist(src) {
		return error.ERROR_NOT_EXIST_IMAGE
	}

	toPath := config.AppSetting.StorageRootPath + config.FileSetting.ImageMovePath + img.Path
	if err := file.MoveFile(src, toPath); err != nil {
		return error.ERROR_DEL_IMAGE_FAIL
	}
	return error.SUCCESS
}
