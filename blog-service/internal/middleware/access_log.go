package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (writer AccessLogWriter) Write(content []byte) (int, error) {
	n, err := writer.body.Write(content)
	if err != nil {
		return n, err
	}
	return writer.ResponseWriter.Write(content)
}

func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: context.Writer}
		context.Writer = bodyWriter
		beginTime := time.Now().Unix()
		context.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  context.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		desc := "access log: method: %s, status_code: %d, " +
			"begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Info(context, desc,
			context.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
