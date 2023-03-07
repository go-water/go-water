package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
	"github.com/gomarkdown/markdown"
	"html/template"
	"io/ioutil"
)

type GetArticleRequest struct {
	UrlID string `json:"url_id"`
}

type GetArticleService struct {
	*water.ServerBase
}

func (srv *GetArticleService) Handle(ctx context.Context, req *GetArticleRequest) (interface{}, error) {
	article, err := model.GetArticle(model.DbMap, req.UrlID)
	if err != nil {
		return nil, err
	}

	mdBytes, err := ioutil.ReadFile(fmt.Sprintf("./content/%s.md", article.UrlID))
	if err != nil {
		return nil, err
	}

	article.Body = template.HTML(markdown.ToHTML(mdBytes, nil, nil))

	return article, nil
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

func (srv *GetArticleService) GetRequest() interface{} {
	return new(GetArticleRequest)
}
