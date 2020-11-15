package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/pkg/errcode"
	"net/http"
)

type Response struct {
	context *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(context *gin.Context) *Response {
	return &Response{
		context: context,
	}
}

func (response * Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	response.context.JSON(http.StatusOK, data)
}

func (response *Response) ToResponseList(list interface{}, totalRows int) {
	response.context.JSON(http.StatusOK, gin.H {
		"list": list,
		"pager": Pager{
			Page: GetPage(response.context),
			PageSize: GetPageSize(response.context),
			TotalRows: totalRows,
		},
	})
}

func (response *Response) ToErrorResponse(err *errcode.Error) {
	res := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		res["details"] = details
	}
	response.context.JSON(err.StatusCode(), res)
}
