package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/service"
	"html/template"
	"net/http"
)

func (h *Handlers) Index(ctx *gin.Context) {
	req := new(service.IndexRequest)
	resp, err := h.index.ServerWater(ctx, req)
	if err != nil {
		h.index.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.([]byte); ok {
		ctx.HTML(http.StatusOK, "index", gin.H{"body": template.HTML(result), "title": "go-water 官方网站"})
	}
}

func (h *Handlers) ListDoc(ctx *gin.Context) {
	req := new(service.ListDocRequest)
	req.Kind = model.ArticleKindDoc
	resp, err := h.listDoc.ServerWater(ctx, req)
	if err != nil {
		h.listDoc.GetLogger().Error(err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "articles", gin.H{"body": resp, "title": "文档"})
}

func (h *Handlers) ListArticle(ctx *gin.Context) {
	req := new(service.ListArticleRequest)
	req.Kind = model.ArticleKindTech
	resp, err := h.listArticle.ServerWater(ctx, req)
	if err != nil {
		h.listArticle.GetLogger().Error(err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "articles", gin.H{"body": resp, "title": "技术文章"})
}

func (h *Handlers) GetArticle(ctx *gin.Context) {
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

	ctx.HTML(http.StatusOK, "detail", gin.H{"body": resp, "title": title})
}
