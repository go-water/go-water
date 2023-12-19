package service

import (
	"bytes"
	"context"
	"github.com/go-water/water"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"os"
)

type IndexRequest struct{}

type IndexService struct {
	*water.ServerBase
}

func (srv *IndexService) Handle(ctx context.Context, req *IndexRequest) (interface{}, error) {
	mdBytes, err := os.ReadFile("./content/index.md")
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

	return buf.Bytes(), nil
}
