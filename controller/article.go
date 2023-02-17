package controller

import (
	"context"
	"github.com/go-water/go-water/service"
	"github.com/kataras/iris/v12"
)

func (h *Handlers) Index(ctx iris.Context) {
	req := new(service.ListArticleRequest)
	resp, err := h.listArticle.ServerWater(context.Background(), req)
	if err != nil {
		ctx.EndRequest()
	}

	ctx.ViewData("title", "爱斯园 - Golang 学习网站")
	ctx.ViewData("body", resp)
	ctx.View("articles.html")
}

func (h *Handlers) ListArticle(ctx iris.Context) {
	req := new(service.ListArticleRequest)
	resp, err := h.listArticle.ServerWater(context.Background(), req)
	if err != nil {
		ctx.EndRequest()
	}

	ctx.ViewData("title", "爱斯园 - Golang 学习网站")
	ctx.ViewData("body", resp)
	ctx.View("articles.html")
}

func (h *Handlers) GetArticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	req := new(service.GetArticleRequest)
	req.UrlID = id
	resp, err := h.getArticle.ServerWater(context.Background(), req)
	if err != nil {
		ctx.EndRequest()
	}

	ctx.ViewData("body", resp)
	ctx.View("detail.html")
}
