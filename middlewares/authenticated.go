package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ronytampubolon/miniwallet/utils"
)

func AuthenticatedUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := utils.VerifyTokenHeader(ctx, "JWT_SECRET")

		if err != nil {
			utils.UnauthorizedError(ctx)
		} else {
			ctx.Set("user", token.Claims)
			ctx.Next()
		}

	}
}
