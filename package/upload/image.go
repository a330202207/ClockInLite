package upload

import (
	"ClockInLite/config"
	"ClockInLite/package/file"
	"ClockInLite/util"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func GetImageFullUrl(name string) string {
	return config.FileSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

//获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	str := string(byte(r.Intn(26)))
	fileName = util.EncodeMD5(str + fileName)

	return fileName + ext
}

//获取图片完整访问URL
func GetImagePath() string {
	return config.FileSetting.ImageSavePath
}

//获取图片完整路径
func GetImageFullPath() string {
	return config.AppSetting.StorageRootPath + GetImagePath()
}

//检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range config.FileSetting.ImageAllowExt {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

//检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}
	return size <= config.FileSetting.ImageMaxSize
}

//检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
