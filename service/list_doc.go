package service

import (
	"context"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
)

type ListDocRequest struct {
	Kind int `json:"kind"`
}

type ListDocService struct {
	*water.ServerBase
}

func (srv *ListDocService) Handle(ctx context.Context, req *ListDocRequest) (interface{}, error) {
	result, err := model.ListArticles(model.DbMap, req.Kind)
	if err != nil {
		return nil, err
	}

	return result, nil
}
