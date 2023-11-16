package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
	"github.com/go-water/water/endpoint"
	"html/template"
	"time"
)

type UpdatePostRequest struct {
	UrlID       string        `form:"url_id"`
	Title       string        `form:"title"`
	Icon        string        `form:"icon"`
	Kind        int           `form:"kind"`
	Brief       string        `form:"brief"`
	Body        template.HTML `form:"body"`
	UpdatedTime time.Time     `json:"UpdatedTime"`
}

type UpdatePostService struct {
	*water.ServerBase
}

func (srv *UpdatePostService) Handle(ctx context.Context, req *UpdatePostRequest) (interface{}, error) {
	err := model.UpdateArticle(model.DbMap, req.UrlID, req.Title, req.Icon, req.Brief, req.Kind, req.Body, time.Now())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (srv *UpdatePostService) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*UpdatePostRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}
