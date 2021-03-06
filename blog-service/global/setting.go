package global

import (
	"github.com/go-microservices/blog-service/pkg/logger"
	"github.com/go-microservices/blog-service/pkg/settting"
)

var (
	ServerSetting   *settting.ServerSettingS
	AppSetting      *settting.AppSettingS
	DatabaseSetting *settting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *settting.JWTSettingS
	EmailSettingS   *settting.EmailSettingS
)
