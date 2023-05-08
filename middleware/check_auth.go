package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-water/go-water/utils"
	"github.com/go-water/water"
)

func CheckAuth(ctx *gin.Context) {
	token, _ := ctx.Cookie(utils.AuthorizationKey)
	if len(token) == 0 {
		ctx.Abort()
		return
	}

	ctx.Request.Header.Set(utils.AuthorizationKey, utils.BearerKey+" "+token)
	userUUID, _, err := water.ParseAndValid(ctx.Request, utils.RsaPublicKeyPath)
	if err != nil {
		ctx.Abort()
		return
	}

	ctx.Set("uuid", userUUID)
}
