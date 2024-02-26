package service

import (
	"context"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
	"github.com/jinzhu/copier"
	"html/template"
)

type AddPostRequest struct {
	UrlID string        `form:"url_id" binding:"required"`
	Title string        `form:"title" binding:"required"`
	Brief string        `form:"brief" binding:"required"`
	Body  template.HTML `form:"body" binding:"required"`
}

type AddPostService struct {
	*water.ServerBase
}

func (srv *AddPostService) Handle(ctx context.Context, req *AddPostRequest) (interface{}, error) {
	article := new(model.Article)
	err := copier.Copy(article, req)
	if err != nil {
		return nil, err
	}

	err = article.Insert(model.DbMap)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
