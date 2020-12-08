package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/internal/service"
	"github.com/go-microservices/blog-service/pkg/app"
	"github.com/go-microservices/blog-service/pkg/errcode"
)

func GetAuth(context *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorFormat(context, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(context.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.ErrorFormat(context, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.ErrorFormat(context, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
