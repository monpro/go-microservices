package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/convert"
)

func GetPage(context *gin.Context) int {
	page := convert.StrTo(context.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(context *gin.Context) int {
	pageSize := convert.StrTo(context.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}

