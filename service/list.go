package service

import (
	"context"
	"errors"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
	"github.com/go-water/water/endpoint"
)

type ListRequest struct {
	PageNo   int `json:"page_no"`
	PageSize int `json:"page_size"`
}

type ListService struct {
	*water.ServerBase
}

func (srv *ListService) Handle(ctx context.Context, req *ListRequest) (interface{}, error) {
	list, err := model.List(model.DbMap)

	return list, err
}

func (srv *ListService) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*ListRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}
