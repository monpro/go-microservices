package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/pkg/app"
	"github.com/go-microservices/blog-service/pkg/errcode"
)

type Article struct {}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context)    {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (t Article) List(c *gin.Context)   {}
func (t Article) Create(c *gin.Context) {}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}

