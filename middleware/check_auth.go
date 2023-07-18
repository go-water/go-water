package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-water/go-water/utils"
	"github.com/go-water/water"
	"net/http"
)

func CheckAuth(ctx *gin.Context) {
	token, _ := ctx.Cookie(utils.AuthorizationKey)
	if len(token) == 0 {
		ctx.Abort()
		ctx.Redirect(http.StatusFound, "/login")
	}

	ctx.Request.Header.Set(utils.AuthorizationKey, utils.BearerKey+" "+token)
	userUUID, _, err := water.ParseFromRequest(ctx.Request, utils.RsaPublicKeyPath)
	if err != nil {
		ctx.Abort()
		return
	}

	ctx.Set("uuid", userUUID)
}
