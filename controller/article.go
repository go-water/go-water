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

	ctx.HTML(http.StatusOK, "doc", water.H{"body": resp, "title": "文档"})
}

func (h *Handlers) GetDoc(ctx *water.Context) {
	id := ctx.Param("id")
	req := new(service.GetDocRequest)
	req.UrlID = id
	resp, err := h.getDoc.ServerWater(ctx, req)
	if err != nil {
		h.getDoc.GetLogger().Error(err.Error())
		return
	}

	title := ""
	if article, ok := resp.(*service.GetDocResponse); ok {
		title = article.Title
	}

	ctx.HTML(http.StatusOK, "doc", water.H{"body": resp, "title": title})
}
