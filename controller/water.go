package controller

import (
	"github.com/go-water/go-water/service"
	"github.com/go-water/water"
	"net/http"
)

func (h *Handlers) Reward(ctx *water.Context) {
	req := new(service.RewardRequest)
	resp, err := h.reward.ServerWater(ctx, req)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.([][]string); ok {
		ctx.HTML(http.StatusOK, "reward", water.H{"body": result, "title": "打赏站长"})
	}
}

func (h *Handlers) About(ctx *water.Context) {
	ctx.HTML(http.StatusOK, "about", water.H{"title": "关于网站"})
}
