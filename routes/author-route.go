package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/middlewares"
	"guptalibrary.com/models"
)

func GroupAuthorRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AuthMiddleware()).POST("/", CreateAuthor)
	router.Use(middlewares.AuthMiddleware()).PATCH("/:id", UpdateAuthor)
	router.Use(middlewares.AuthMiddleware()).GET("/", GetAuthors)
}

func CreateAuthor(ctx *gin.Context) {

	author := &models.Author{}
	ctx.ShouldBindBodyWithJSON(author)
	id, dbError := handlers.CreateAuthor(author)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}
	ctx.JSON(201, gin.H{"id": id, "message": "Author created successfully"})
}

func UpdateAuthor(ctx *gin.Context) {

	author := &models.Author{}
	ctx.ShouldBindBodyWithJSON(author)

	id := ctx.Param("id")
	authorId, parseError := strconv.ParseInt(id, 10, 64)
	if parseError != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	generatedId, dbError := handlers.UpdateAuthor(int(authorId), author)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": generatedId, "message": "Author updated successfully"})
}

func GetAuthors(ctx *gin.Context) {
	authors, dbError := handlers.GetAuthors()
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}
	ctx.JSON(200, authors)
}
