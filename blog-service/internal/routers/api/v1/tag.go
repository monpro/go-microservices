package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/internal/service"
	"github.com/go-microservices/blog-service/pkg/app"
	"github.com/go-microservices/blog-service/pkg/convert"
	"github.com/go-microservices/blog-service/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

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
func (t Tag) List(context *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorFormat(context, "app bindAndValid errors %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(context.Request.Context())
	pager := app.Pager{Page: app.GetPage(context), PageSize: app.GetPageSize(context)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.ErrorFormat(context, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.ErrorFormat(context, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}

//@Summary create a tag
//@Produce json
//@Param name query string false "tagName" maxlength(100)
//@Param name query int false "status" Enums(0, 1) default(1)
//@Param created_by body string false "creator" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorFormat(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.ErrorFormat(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return

}

//@Summary update  atag
//@Produce json
//@Param id path int true "id for tag"
//@Param name body string false "tag name" minlength(3) maxlength(100)
//@Param state body int false "state" Enums(0, 1) default(1)
//@Param modified_by body string true "modifier" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorFormat(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.ErrorFormat(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//@Summary delete a tag
//@Produce json
//@Param id path int true "id for tag"
//@Success 200 {object} string "Success"
//@Failure 400 {object} errcode.Error "Failure"
//@Failure 500 {object} errcode.Error "Internal Server Error"
//@Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorFormat(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.ErrorFormat(c, "svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return

}
