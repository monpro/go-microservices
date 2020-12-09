package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-microservices/blog-service/docs"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/internal/middleware"
	"github.com/go-microservices/blog-service/internal/routers/api"
	v1 "github.com/go-microservices/blog-service/internal/routers/api/v1"
	"github.com/go-microservices/blog-service/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var limiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.BucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	router := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	} else {
		router.Use(middleware.AccessLog())
		router.Use(middleware.Recovery())
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Translation())
	router.Use(middleware.Tracing())

	router.Use(middleware.RateLimiter(limiters))
	router.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	router.POST("/auth", api.GetAuth)
	router.POST("/upload/file", upload.UploadFile)
	router.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiV1 := router.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}
	return router
}
