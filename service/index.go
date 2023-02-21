package service

import (
	"context"
	"errors"
	"github.com/go-water/water"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
)

type IndexRequest struct{}

type IndexService struct {
	*water.ServerBase
}

func (srv *IndexService) Handle(ctx context.Context, req *IndexRequest) (interface{}, error) {
	mdBytes, err := ioutil.ReadFile("./content/index.md")
	if err != nil {
		return nil, err
	}

	html := markdown.ToHTML(mdBytes, nil, nil)

	return html, nil
}

func (srv *IndexService) Endpoint() water.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if r, ok := req.(*IndexRequest); ok {
			return srv.Handle(ctx, r)
		} else {
			return nil, errors.New("request type error")
		}
	}
}

func (srv *IndexService) Name() string {
	return srv.ServerBase.Name(srv)
}
