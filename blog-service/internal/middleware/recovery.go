package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/app"
	"github.com/go-microservices/blog-service/pkg/email"
	"github.com/go-microservices/blog-service/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSettingS.Host,
		Port:     global.EmailSettingS.Port,
		IsSSL:    global.EmailSettingS.IsSSL,
		UserName: global.EmailSettingS.UserName,
		Password: global.EmailSettingS.Password,
		From:     global.EmailSettingS.From,
	})
	return func(context *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				global.Logger.WithCallersFrames().ErrorFormat(context, "panic recover err: %v", err)

				err := defaultMailer.SendMail(
					global.EmailSettingS.To,
					fmt.Sprintf("exception at: %d", time.Now().Unix()),
					fmt.Sprintf("err: %v", err),
				)
				if err != nil {
					global.Logger.PanicFormat(context, "mail.SendMail err: %v", err)
				}
				app.NewResponse(context).ToErrorResponse(errcode.ServerError)
				context.Abort()
			}
		}()
		context.Next()
	}
}
