package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
)

type ArticleRequest struct {
	Size int `json:"size"`
}

type ArticleService struct {
	*water.ServerBase
}

func (srv *ArticleService) Handle(ctx context.Context, req *ArticleRequest) (interface{}, error) {
	result, err := model.ListArticles(model.DbMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *ArticleService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*ArticleRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}

func (srv *ArticleService) Name() string {
	return srv.ServerBase.Name(srv)
}
