package v1

import "github.com/gin-gonic/gin"

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context)    {}
//@Summary get a list of tags
//@Produce json
//@Param name query string false "tagName" maxlength(100)
//@Param name query int false "status" Enums(0, 1) default(1)
//@Param page query int false "page"
//@Param page_size query int false "page_size"
//@Success 200 {object} model.TagSwapper "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context)   {}
//@Summary create a tag
//@Produce json
//@Param name query string false "tagName" maxlength(100)
//@Param name query int false "status" Enums(0, 1) default(1)
//@Param created_by body string false "creator" minlength(3) maxlength(100)
//@Success 200 {object} model.TagSwapper "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}
//@Summary update  atag
//@Produce json
//@Param id path int true "id for tag"
//@Param name body string false "tag name" minlength(3) maxlength(100)
//@Param state body int false "state" Enums(0, 1) default(1)
//@Param modified_by body string true "modifier" minlength(3) maxlength(100)
//@Success 200 {object} model.TagSwapper "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}
//@Summary delete a tag
//@Produce json
//@Param id path int true "id for tag"
//@Success 200 {object} model.TagSwapper "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
