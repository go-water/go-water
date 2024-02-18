package main

import (
	"github.com/go-water/go-water/controller"
	"github.com/go-water/go-water/helpers"
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/router"
)

func main() {
	router.Start()
}

func init() {
	helpers.InitConfig()
	model.InitDB()
	controller.InitService()
}
