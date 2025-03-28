package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/utils"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var authorization string = ctx.GetHeader("Authorization")

		if authorization == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		token, found := strings.CutPrefix(authorization, "Bearer ")
		if !found {
			ctx.JSON(401, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		isValid, err := utils.ValidateToken(token)

		if err != nil || !isValid {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("username", utils.GetUsernameFromToken(token))
		ctx.Next()
	}
}
