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
	loginPost   water.Handler
	list        water.Handler
	addPost     water.Handler
	update      water.Handler
	updatePost  water.Handler
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
		loginPost:   water.NewHandler(&service.LoginPostService{ServerBase: &water.ServerBase{}}, option),
		list:        water.NewHandler(&service.ListService{ServerBase: &water.ServerBase{}}, option),
		addPost:     water.NewHandler(&service.AddPostService{ServerBase: &water.ServerBase{}}, option),
		update:      water.NewHandler(&service.UpdateService{ServerBase: &water.ServerBase{}}, option),
		updatePost:  water.NewHandler(&service.UpdatePostService{ServerBase: &water.ServerBase{}}, option),
	}
}
