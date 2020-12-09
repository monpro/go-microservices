package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/app"
	"github.com/go-microservices/blog-service/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				global.Logger.WithCallersFrames().ErrorFormat(context, "panic recover err: %v", err)
				app.NewResponse(context).ToErrorResponse(errcode.ServerError)
				context.Abort()
			}
		}()
		context.Next()
	}
}
