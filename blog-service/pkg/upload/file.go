package upload

import (
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
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

func GetServerUrl() string {
	return global.AppSetting.UploadServerUrl
}

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func CheckSavePath(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func CheckContainExt(fileType FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch fileType {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(fileType FileType, file multipart.File) bool {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return false
	}
	size := len(content)
	switch fileType {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CheckPermission(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

func CreateSavePath(path string, permission os.FileMode) error {
	err := os.MkdirAll(path, permission)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
