package service

import (
	"context"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
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
