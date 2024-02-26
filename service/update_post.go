package service

import (
	"context"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
	"html/template"
	"time"
)

type UpdatePostRequest struct {
	UrlID       string        `form:"url_id"`
	Title       string        `form:"title"`
	Brief       string        `form:"brief"`
	Body        template.HTML `form:"body"`
	UpdatedTime time.Time     `json:"UpdatedTime"`
}

type UpdatePostService struct {
	*water.ServerBase
}

func (srv *UpdatePostService) Handle(ctx context.Context, req *UpdatePostRequest) (interface{}, error) {
	err := model.UpdateArticle(model.DbMap, req.UrlID, req.Title, req.Brief, req.Body, time.Now())
	if err != nil {
		return nil, err
	}

	return nil, nil
}
