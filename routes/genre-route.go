package routes

import (
	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/models"
)

func GroupGenreRoutes(router *gin.RouterGroup) {

	router.POST("/", CreateGenre)
}

func CreateGenre(ctx *gin.Context) {

	genre := &models.Genre{}
	ctx.ShouldBindBodyWithJSON(genre)

	id, err := handlers.CreateGenre(genre)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": id, "message": "genre created successfully"})
}
