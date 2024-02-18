package middleware

import (
	"github.com/go-water/go-water/utils"
	"github.com/go-water/water"
	"net/http"
)

func CheckAuth(handler water.HandlerFunc) water.HandlerFunc {
	return func(ctx *water.Context) {
		token, _ := ctx.Cookie(utils.AuthorizationKey)
		if len(token) == 0 {
			ctx.Redirect(http.StatusFound, "/login")
		}

		ctx.Request.Header.Set(utils.AuthorizationKey, utils.BearerKey+" "+token)
		userUUID, issuer, _, err := water.ParseFromRequest(ctx.Request, utils.RsaPublicKeyPath)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/login")
		}

		ctx.Set("uuid", userUUID)
		ctx.Set("issuer", issuer)

		handler(ctx)
	}
}
