package routes

import (
	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/models"
)

func GroupBookRoutes(router *gin.RouterGroup) {

	router.POST("/", CreateBook)
}

func CreateBook(ctx *gin.Context) {

	newBook := &models.Book{}
	ctx.ShouldBindBodyWithJSON(newBook)

	generatedId, dbError := handlers.CreateBook(newBook)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": generatedId, "message": "Book created successfully"})
}
