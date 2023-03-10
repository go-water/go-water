package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
)

type ListArticleRequest struct {
	Kind int `json:"kind"`
}

type ListArticleService struct {
	*water.ServerBase
}

func (srv *ListArticleService) Handle(ctx context.Context, req *ListArticleRequest) (interface{}, error) {
	result, err := model.ListArticles(model.DbMap, req.Kind)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *ListArticleService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*ListArticleRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}
