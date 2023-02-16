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
	article water.Handler
}

func NewService() *Handlers {
	conf := &water.Config{Encoding: "console", Level: zap.InfoLevel}
	option := water.ServerConfig(conf)
	return &Handlers{
		article: water.NewHandler(&service.ArticleService{ServerBase: &water.ServerBase{}}, option),
	}
}
