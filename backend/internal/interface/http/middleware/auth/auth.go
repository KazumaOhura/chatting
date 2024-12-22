package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header
		token := header.Get("Authorization")

		// ヘッダーにトークンが無い
		if token == "" {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		// Parse token

	}
}
