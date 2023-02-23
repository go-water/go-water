package controller

import (
	"context"
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/service"
	"github.com/kataras/iris/v12"
	"html/template"
)

func (h *Handlers) Index(ctx iris.Context) {
	req := new(service.IndexRequest)
	resp, err := h.index.ServerWater(context.Background(), req)
	if err == nil && resp != nil {
		if result, ok := resp.([]byte); ok {
			ctx.ViewData("body", template.HTML(result))
		}
	}

	ctx.ViewData("title", "爱斯园 - go-water 官方网站")
	ctx.View("index.html")
}

func (h *Handlers) ListDoc(ctx iris.Context) {
	req := new(service.ListDocRequest)
	req.Kind = model.ArticleKindDoc
	resp, err := h.listDoc.ServerWater(context.Background(), req)
	if err == nil {
		ctx.ViewData("body", resp)
	}

	ctx.ViewData("title", "文档 - 爱斯园")
	ctx.View("articles.html")
}

func (h *Handlers) ListArticle(ctx iris.Context) {
	req := new(service.ListArticleRequest)
	req.Kind = model.ArticleKindTech
	resp, err := h.listArticle.ServerWater(context.Background(), req)
	if err == nil {
		ctx.ViewData("body", resp)
	}

	ctx.ViewData("title", "技术文章 - 爱斯园")
	ctx.View("articles.html")
}

func (h *Handlers) GetArticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	req := new(service.GetArticleRequest)
	req.UrlID = id
	resp, err := h.getArticle.ServerWater(context.Background(), req)
	if err == nil {
		ctx.ViewData("body", resp)
	}

	ctx.View("detail.html")
}
