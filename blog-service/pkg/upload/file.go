package upload

import (
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/util"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}
