package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/models"
)

func GroupBookRoutes(router *gin.RouterGroup) {

	router.POST("/", CreateBook)
	router.PATCH("/:id", UpdateBook)
	router.GET("/", GetBooks)
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

func UpdateBook(ctx *gin.Context) {

	param := ctx.Param("id")
	if param == "" {
		ctx.JSON(400, gin.H{"error": "Book ID is required"})
		return
	}
	bookId, parseError := strconv.ParseInt(param, 10, 64)
	if parseError != nil {
		ctx.JSON(400, gin.H{"error": "Invalid Book ID"})
		return
	}
	book := &models.Book{}
	ctx.ShouldBindBodyWithJSON(book)

	generatedId, dbError := handlers.UpdateBook(int(bookId), book)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": generatedId, "message": "Book updated successfully"})
}

func GetBooks(ctx *gin.Context) {

	books, dbError := handlers.GetBooks()
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"books": books})
}
