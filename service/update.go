package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
)

type UpdateRequest struct {
	UrlID string `json:"url_id"`
}

type UpdateService struct {
	*water.ServerBase
}

func (srv *UpdateService) Handle(ctx context.Context, req *UpdateRequest) (interface{}, error) {
	article, err := model.GetArticle(model.DbMap, req.UrlID)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (srv *UpdateService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*UpdateRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}
