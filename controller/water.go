package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-water/go-water/service"
	"net/http"
)

func (h *Handlers) Reward(ctx *gin.Context) {
	req := new(service.RewardRequest)
	resp, err := h.reward.ServerWater(ctx, req)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.([][]string); ok {
		ctx.HTML(http.StatusOK, "reward", gin.H{"body": result, "title": "打赏站长"})
	}
}

func (h *Handlers) About(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "about", gin.H{"title": "关于网站"})
}
