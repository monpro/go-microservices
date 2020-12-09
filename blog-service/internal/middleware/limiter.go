package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/pkg/app"
	"github.com/go-microservices/blog-service/pkg/errcode"
	"github.com/go-microservices/blog-service/pkg/limiter"
)

func RateLimiter(l limiter.ILimiter) gin.HandlerFunc {
	return func(context *gin.Context) {
		key := l.Key(context)
		bucket, valid := l.GetBucket(key)
		if valid {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(context)
				response.ToErrorResponse(errcode.TooManyRequests)
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
