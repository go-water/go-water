package controller

import "github.com/kataras/iris/v12"

func (h *Handlers) Reward(ctx iris.Context) {
	ctx.ViewData("title", "打赏站长")
	ctx.View("reward.html")
}

func (h *Handlers) About(ctx iris.Context) {
	ctx.ViewData("title", "关于网站")
	ctx.ViewData("body", 88888)
	ctx.View("about.html")
}

func (h *Handlers) Link(ctx iris.Context) {
	ctx.ViewData("title", "友情链接")
	ctx.View("link.html")
}
