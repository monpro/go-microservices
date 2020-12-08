package service

import (
	"errors"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType,
	file multipart.File,
	header *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(header.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file type is not supported")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("maximum size is exceeded")
	}
	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("no permissions")
	}
	path := uploadSavePath + "/" + fileName
	err := upload.SaveFile(header, path)
	if err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
