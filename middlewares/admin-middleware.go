package middlewares

import (
	"github.com/gin-gonic/gin"
	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func AdminMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		username, doesExists := ctx.Get("username")
		if !doesExists {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "username not available in token!"})
			return
		}

		existingUser := &models.User{}

		configs.DB.
			Where("username = ?", username).
			Joins("JOIN user_roles on user_roles.user_id = users.id").
			Preload("Roles").
			Find(existingUser)

		var roles []models.Role = existingUser.Roles

		if len(roles) == 0 {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "no roles are assigned to the user"})
			return
		}

		var isAdmin bool = false
		for _, role := range roles {
			if role.Name == "ADMIN" {
				isAdmin = true
				break
			}
		}

		if !isAdmin {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "user is not eligible to perform the action"})
			return
		}

		ctx.Next()
	}
}
