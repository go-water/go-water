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
	index       water.Handler
	reward      water.Handler
	listDoc     water.Handler
	listArticle water.Handler
	getArticle  water.Handler
}

func NewService() *Handlers {
	conf := &water.Config{Encoding: "console", Level: zap.InfoLevel}
	option := water.ServerConfig(conf)
	return &Handlers{
		index:       water.NewHandler(&service.IndexService{ServerBase: &water.ServerBase{}}, option),
		reward:      water.NewHandler(&service.RewardService{ServerBase: &water.ServerBase{}}, option),
		listDoc:     water.NewHandler(&service.ListDocService{ServerBase: &water.ServerBase{}}, option),
		listArticle: water.NewHandler(&service.ListArticleService{ServerBase: &water.ServerBase{}}, option),
		getArticle:  water.NewHandler(&service.GetArticleService{ServerBase: &water.ServerBase{}}, option),
	}
}
