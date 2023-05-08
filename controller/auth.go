package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-water/go-water/service"
	"github.com/go-water/go-water/utils"
	"github.com/spf13/viper"
	"net/http"
)

func (h *Handlers) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", gin.H{"title": "登录"})
}

func (h *Handlers) LoginPost(ctx *gin.Context) {
	request := new(service.LoginPostRequest)
	if err := ShouldBind(ctx, request); err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		return
	}

	resp, err := h.loginPost.ServerWater(ctx, request)
	if err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		return
	}

	if result, ok := resp.(string); ok {
		ctx.SetCookie(utils.AuthorizationKey, result, int(utils.AuthTimeout.Seconds()), "/", viper.GetString("service.domain"), false, true)
		ctx.Redirect(http.StatusFound, "/admin/list")
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}
}
