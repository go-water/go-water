package controller

import (
	"github.com/go-water/go-water/service"
	"github.com/go-water/water"
	"github.com/sony/gobreaker"
	"time"
)

var (
	H *Handlers
)

type Handlers struct {
	index      water.Handler
	reward     water.Handler
	listDoc    water.Handler
	getDoc     water.Handler
	loginPost  water.Handler
	list       water.Handler
	addPost    water.Handler
	update     water.Handler
	updatePost water.Handler
}

func NewService() *Handlers {
	option := water.ServerLimiter(time.Second, 10000)
	return &Handlers{
		index:      water.NewHandler(&service.IndexService{ServerBase: &water.ServerBase{}}, option),
		reward:     water.NewHandler(&service.RewardService{ServerBase: &water.ServerBase{}}, option),
		listDoc:    water.NewHandler(&service.ListDocService{ServerBase: &water.ServerBase{}}, option),
		getDoc:     water.NewHandler(&service.GetDocService{ServerBase: &water.ServerBase{}}, option),
		loginPost:  water.NewHandler(&service.LoginPostService{ServerBase: &water.ServerBase{}}, option, water.ServerBreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))),
		list:       water.NewHandler(&service.ListService{ServerBase: &water.ServerBase{}}, option),
		addPost:    water.NewHandler(&service.AddPostService{ServerBase: &water.ServerBase{}}, option),
		update:     water.NewHandler(&service.UpdateService{ServerBase: &water.ServerBase{}}, option),
		updatePost: water.NewHandler(&service.UpdatePostService{ServerBase: &water.ServerBase{}}, option),
	}
}

func InitService() {
	H = NewService()
}
