package middleware

import (
	"fmt"
	"github.com/go-water/water"
	"time"
)

func TraceApi(handlerFunc water.HandlerFunc) water.HandlerFunc {
	return func(ctx *water.Context) {
		start := time.Now()
		defer func() {
			msg := fmt.Sprintf("[WATER] %v | %15s | %13v | %-7s %s",
				time.Now().Format("2006/01/02 - 15:04:05"),
				ctx.ClientIP(),
				time.Since(start),
				ctx.Request.Method,
				ctx.Request.URL.Path,
			)

			fmt.Println(msg)
		}()

		handlerFunc(ctx)
	}
}
