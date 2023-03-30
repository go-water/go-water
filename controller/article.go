package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/service"
	"html/template"
	"net/http"
)

func (h *Handlers) Index(ctx *gin.Context) {
	//GetRequest 内部使用 pool，可以减少内存分配，无需每次初始化对象
	// 效果下如下直接 new 一个对象一样
	//req := new(service.IndexRequest)
	req := h.index.GetRequest()
	resp, err := h.index.ServerWater(context.Background(), req)
	if err != nil {
		h.index.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.([]byte); ok {
		ctx.HTML(http.StatusOK, "index", gin.H{"body": template.HTML(result), "title": "go-water 官方网站"})
	}
}

func (h *Handlers) ListDoc(ctx *gin.Context) {
	req := h.listDoc.GetRequest().(*service.ListDocRequest)
	req.Kind = model.ArticleKindDoc
	resp, err := h.listDoc.ServerWater(context.Background(), req)
	if err != nil {
		h.listDoc.GetLogger().Error(err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "articles", gin.H{"body": resp, "title": "文档"})
}

func (h *Handlers) ListArticle(ctx *gin.Context) {
	req := new(service.ListArticleRequest)
	req.Kind = model.ArticleKindTech
	resp, err := h.listArticle.ServerWater(context.Background(), req)
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
	resp, err := h.getArticle.ServerWater(context.Background(), req)
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
