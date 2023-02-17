package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
)

type GetArticleRequest struct {
	UrlID string `json:"url_id"`
}

type GetArticleService struct {
	*water.ServerBase
}

func (srv *GetArticleService) Handle(ctx context.Context, req *GetArticleRequest) (interface{}, error) {
	result, err := model.GetArticle(model.DbMap, req.UrlID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *GetArticleService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*GetArticleRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}

func (srv *GetArticleService) Name() string {
	return srv.ServerBase.Name(srv)
}
