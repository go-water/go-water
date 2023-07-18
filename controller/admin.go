package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-water/go-water/model"
	"github.com/go-water/go-water/service"
	"net/http"
)

func (h *Handlers) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add", gin.H{"title": "添加文章"})
}

func (h *Handlers) AddPost(ctx *gin.Context) {
	request := new(service.AddPostRequest)
	if err := ShouldBind(ctx, request); err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	_, err := h.addPost.ServerWater(context.Background(), request)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func (h *Handlers) List(ctx *gin.Context) {
	req := new(service.ListRequest)
	resp, err := h.list.ServerWater(ctx, req)
	if err != nil {
		h.listDoc.GetLogger().Error(err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "list", gin.H{"body": resp, "title": "文章列表"})
}

func (h *Handlers) Update(ctx *gin.Context) {
	req := new(service.UpdateRequest)
	req.UrlID = ctx.Param("id")
	resp, err := h.update.ServerWater(ctx, req)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.(*model.Article); ok {
		ctx.HTML(http.StatusOK, "update", gin.H{"body": result, "title": "修改文章"})
	}
}

func (h *Handlers) UpdatePost(ctx *gin.Context) {
	request := new(service.UpdatePostRequest)
	if err := ShouldBind(ctx, request); err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	_, err := h.updatePost.ServerWater(ctx, request)
	if err != nil {
		h.reward.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true})
}
