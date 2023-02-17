package controller

import (
	"github.com/go-water/go-water/service"
	"github.com/go-water/water"
	"go.uber.org/zap"
)

var (
	H *Handlers
)

func init() {
	H = NewService()
}

type Handlers struct {
	listArticle water.Handler
	getArticle  water.Handler
}

func NewService() *Handlers {
	conf := &water.Config{Encoding: "console", Level: zap.InfoLevel}
	option := water.ServerConfig(conf)
	return &Handlers{
		listArticle: water.NewHandler(&service.ListArticleService{ServerBase: &water.ServerBase{}}, option),
		getArticle:  water.NewHandler(&service.GetArticleService{ServerBase: &water.ServerBase{}}, option),
	}
}
