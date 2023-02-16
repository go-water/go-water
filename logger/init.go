package logger

import (
	"github.com/go-water/go-water/helpers"
	"go.uber.org/zap"
)

var (
	L *zap.Logger
)

func InitLogger() {
	zLog := helpers.Config{Encoding: "console"}
	L = zLog.NewLogger()
}
