package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct {}

func NewArticle() Article {
	return Article{}
}
//@Summary get an article
//@Produce json
//@Param id path int true "id for the article"
//@Success 200 {object} model.Article "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context)    {
}
//@Summary get a list of articles
//@Produce json
//@Param name query string false "articleName"
//@Param tag_id query int false "tagId"
//@Param state query int false "state"
//@Param page query int false "page"
//@Param page_size query int false "page_size"
//@Success 200 {object} model.ArticleSwagger "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/articles [get]
func (t Article) List(c *gin.Context)   {}
//@Summary get a list of articles
//@Produce json
//@Param tag_id body string true "tag_id"
//@Param title body string true "title"
//@Param desc body string false "description for the article"
//@Param cover_image_url body string true "cover_image_url"
//@Param content body string true "content"
//@Param created_by body int true "creator"
//@Param state body int false "state"
//@Success 200 {object} model.ArticleSwagger "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags [get]
func (t Article) Create(c *gin.Context) {}
//@Summary get a list of articles
//@Produce json
//@Param tag_id body string true "tag_id"
//@Param title body string true "title"
//@Param desc body string false "description for the article"
//@Param cover_image_url body string true "cover_image_url"
//@Param content body string true "content"
//@Param modified_by body int true "modifier"
//@Success 200 {object} model.ArticleSwagger "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {}
//@Summary get a list of articles
//@Produce json
////@Param id path int true "id for the article"
//@Success 200 {object} string "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {}

