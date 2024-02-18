package controller

import (
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/service"
	"github.com/go-water/water"
	"html/template"
	"net/http"
)

func (h *Handlers) Index(ctx *water.Context) {
	req := new(service.IndexRequest)
	resp, err := h.index.ServerWater(ctx, req)
	if err != nil {
		h.index.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.([]byte); ok {
		ctx.HTML(http.StatusOK, "index", water.H{"body": template.HTML(result), "title": "go-water 官方网站"})
	}
}

func (h *Handlers) ListDoc(ctx *water.Context) {
	req := new(service.ListDocRequest)
	req.Kind = model.ArticleKindDoc
	resp, err := h.listDoc.ServerWater(ctx, req)
	if err != nil {
		h.listDoc.GetLogger().Error(err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "articles", water.H{"body": resp, "title": "文档"})
}

func (h *Handlers) ListArticle(ctx *water.Context) {
	req := new(service.ListArticleRequest)
	req.Kind = model.ArticleKindTech
	resp, err := h.listArticle.ServerWater(ctx, req)
	if err != nil {
		h.listArticle.GetLogger().Error(err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "articles", water.H{"body": resp, "title": "技术文章"})
}

func (h *Handlers) GetArticle(ctx *water.Context) {
	id := ctx.Param("id")
	req := new(service.GetArticleRequest)
	req.UrlID = id
	resp, err := h.getArticle.ServerWater(ctx, req)
	if err != nil {
		h.getArticle.GetLogger().Error(err.Error())
		return
	}

	title := ""
	if article, ok := resp.(*model.Article); ok {
		title = article.Title
	}

	ctx.HTML(http.StatusOK, "detail", water.H{"body": resp, "title": title})
}
