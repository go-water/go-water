package controller

import (
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/service"
	"github.com/go-water/water"
	"net/http"
)

func (h *Handlers) Add(ctx *water.Context) {
	ctx.HTML(http.StatusOK, "add", water.H{"title": "添加文章"})
}

func (h *Handlers) AddPost(ctx *water.Context) {
	request, err := water.BindJSON[service.AddPostRequest](ctx)
	if err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, water.H{"msg": err.Error()})
		return
	}

	_, err = h.addPost.ServerWater(ctx, request)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, water.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, water.H{"result": true})
}

func (h *Handlers) List(ctx *water.Context) {
	req := new(service.ListRequest)
	resp, err := h.list.ServerWater(ctx, req)
	if err != nil {
		h.listDoc.GetLogger().Error(err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "list", water.H{"body": resp, "title": "文章列表"})
}

func (h *Handlers) Update(ctx *water.Context) {
	req := new(service.UpdateRequest)
	req.UrlID = ctx.Param("id")
	resp, err := h.update.ServerWater(ctx, req)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.(*model.Article); ok {
		ctx.HTML(http.StatusOK, "update", water.H{"body": result, "title": "修改文章"})
	}
}

func (h *Handlers) UpdatePost(ctx *water.Context) {
	request, err := water.BindJSON[service.UpdatePostRequest](ctx)
	if err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, water.H{"msg": err.Error()})
		return
	}

	_, err = h.updatePost.ServerWater(ctx, request)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, water.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, water.H{"result": true})
}
