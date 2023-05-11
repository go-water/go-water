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
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := h.loginPost.ServerWater(ctx, request)
	if err != nil {
		h.loginPost.GetLogger().Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	result := resp.(string)
	ctx.SetCookie(utils.AuthorizationKey, result, int(utils.AuthTimeout.Seconds()), "/", viper.GetString("service.domain"), false, true)
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func (h *Handlers) Logout(ctx *gin.Context) {
	ctx.SetCookie(utils.AuthorizationKey, "", -1, "/", viper.GetString("service.domain"), false, true)
	ctx.Redirect(http.StatusFound, "/")
}
