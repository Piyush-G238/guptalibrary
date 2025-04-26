package routes

import (
	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/middlewares"
	"guptalibrary.com/models"
)

func GroupRoleRoute(router *gin.RouterGroup) {

	router.
		Use(middlewares.AuthenticationMiddleware()).
		Use(middlewares.AdminMiddleware()).
		POST("/", CreateRole)
}

func CreateRole(ctx *gin.Context) {

	role := &models.Role{}
	ctx.ShouldBindBodyWithJSON(role)

	newId, dbError := handlers.CreateRole(role)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": newId, "message": "Role created successfully"})
}
