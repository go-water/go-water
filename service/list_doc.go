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
	"os"
)

type ListDocRequest struct {
	Kind int `json:"kind"`
}

type ListDocResponse struct {
	Body template.HTML    `json:"body"`
	List []*model.Article `json:"list"`
}

type ListDocService struct {
	*water.ServerBase
}

func (srv *ListDocService) Handle(ctx context.Context, req *ListDocRequest) (interface{}, error) {
	mdBytes, err := os.ReadFile("./content/doc.md")
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
	if err := md.Convert(mdBytes, &buf); err != nil {
		return nil, err
	}

	result, err := model.ListArticles(model.DbMap)
	if err != nil {
		return nil, err
	}

	return &ListDocResponse{
		Body: template.HTML(buf.Bytes()),
		List: result,
	}, nil
}
