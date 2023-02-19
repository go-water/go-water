package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
)

type ListDocRequest struct {
	Size int `json:"size"`
}

type ListDocService struct {
	*water.ServerBase
}

func (srv *ListDocService) Handle(ctx context.Context, req *ListDocRequest) (interface{}, error) {
	result, err := model.ListArticles(model.DbMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *ListDocService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*ListDocRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}

func (srv *ListDocService) Name() string {
	return srv.ServerBase.Name(srv)
}
