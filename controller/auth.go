package controller

import (
	"github.com/go-water/go-water/service"
	"github.com/go-water/go-water/utils"
	"github.com/go-water/water"
	"github.com/spf13/viper"
	"net/http"
)

func (h *Handlers) Login(ctx *water.Context) {
	ctx.HTML(http.StatusOK, "login", water.H{"title": "登录"})
}

func (h *Handlers) LoginPost(ctx *water.Context) {
	request, err := water.BindJSON[service.LoginPostRequest](ctx)
	if err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, water.H{"msg": err.Error()})
		return
	}

	resp, err := h.loginPost.ServerWater(ctx, request)
	if err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, water.H{"msg": err.Error()})
		return
	}

	result := resp.(string)
	ctx.SetCookie(utils.AuthorizationKey, result, int(utils.AuthTimeout.Seconds()), "/", viper.GetString("service.domain"), false, true)
	ctx.JSON(http.StatusOK, water.H{"result": true})
}

func (h *Handlers) Logout(ctx *water.Context) {
	ctx.SetCookie(utils.AuthorizationKey, "", -1, "/", viper.GetString("service.domain"), false, true)
	ctx.Redirect(http.StatusFound, "/")
}
