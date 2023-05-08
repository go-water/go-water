package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ShouldBind(ctx *gin.Context, obj interface{}) (err error) {
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
	}()

	if err = ctx.ShouldBind(obj); err != nil {
		return
	}

	return
}
