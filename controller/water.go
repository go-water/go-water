package controller

import "github.com/kataras/iris/v12"

func (h *Handlers) Reward(ctx iris.Context) {
	ctx.ViewData("title", "打赏站长")
	ctx.ViewData("body", 88888)
	ctx.View("reward.html")
}
