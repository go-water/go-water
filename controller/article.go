package controller

import (
	"context"
	"github.com/go-water/go-water/service"
	"github.com/kataras/iris/v12"
)

func (h *Handlers) ListArticle(ctx iris.Context) {
	req := new(service.ListArticleRequest)
	resp, err := h.listArticle.ServerWater(context.Background(), req)
	if err != nil {
		ctx.EndRequest()
	}

	ctx.ViewData("title", "文章列表")
	ctx.ViewData("body", resp)
	ctx.View("articles.html")
}
