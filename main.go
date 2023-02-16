package main

import (
	"github.com/go-water/go-water/helpers"
	"github.com/go-water/go-water/logger"
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/router"
)

func main() {
	logger.InitLogger()
	helpers.InitConfig()
	model.InitDB()
	router.Start()
}
