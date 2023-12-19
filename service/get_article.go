package service

import (
	"bytes"
	"context"
	"github.com/go-water/go-water/model"
	"github.com/go-water/water"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"html/template"
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

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(article.Body), &buf); err != nil {
		return nil, err
	}

	article.Body = template.HTML(buf.Bytes())
	return article, nil
}
