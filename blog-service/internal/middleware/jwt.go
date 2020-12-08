package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/pkg/app"
	"github.com/go-microservices/blog-service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			token   string
			errCode = errcode.Success
		)

		queryToken, exist := context.GetQuery("token")
		if exist {
			token = queryToken
		} else {
			token = context.GetHeader("token")
		}
		if token == "" {
			errCode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					errCode = errcode.UnauthorizedTokenTimeOut
				default:
					errCode = errcode.UnauthorizedTokenError
				}
			}
		}
		if errCode != errcode.Success {
			response := app.NewResponse(context)
			response.ToErrorResponse(errCode)
			context.Abort()
			return
		}
		context.Next()
	}
}
