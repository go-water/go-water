package controller

import (
	"context"
	"github.com/go-water/go-water/service"
	"github.com/kataras/iris/v12"
)

func (h *Handlers) Reward(ctx iris.Context) {
	req := new(service.RewardRequest)
	resp, err := h.reward.ServerWater(context.Background(), req)
	if err == nil && resp != nil {
		if result, ok := resp.([][]string); ok {
			ctx.ViewData("body", result)
		}
	}

	ctx.ViewData("title", "打赏站长 - 爱斯园")
	ctx.View("reward.html")
}

func (h *Handlers) About(ctx iris.Context) {
	ctx.ViewData("title", "关于网站 - 爱斯园")
	ctx.ViewData("body", 88888)
	ctx.View("about.html")
}
