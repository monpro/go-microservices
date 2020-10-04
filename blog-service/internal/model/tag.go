package model

import "github.com/go-microservices/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwapper struct {
	List []*Tag
	Pager *app.Pager
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

func (a Tag) TableName() string {
	return "blog_tag"
}
