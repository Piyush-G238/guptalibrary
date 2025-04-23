package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/middlewares"
	"guptalibrary.com/models"
)

func GroupBookRoutes(router *gin.RouterGroup) {

	router.Use(middlewares.AuthMiddleware()).POST("/", CreateBook)
	router.Use(middlewares.AuthMiddleware()).PATCH("/:id", UpdateBook)
	router.Use(middlewares.AuthMiddleware()).GET("/", GetBooks)
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

	searchValue := ctx.Query("search")

	str1 := ctx.Query("page_number")
	str2 := ctx.Query("page_size")

	var pageNumber, pageSize int64
	if str1 == "" {
		pageNumber = 1
	} else {
		pageNumber, _ = strconv.ParseInt(str1, 10, 64)
	}

	if str2 == "" {
		pageSize = 10
	} else {
		pageSize, _ = strconv.ParseInt(ctx.Query("page_size"), 10, 64)
	}

	str3 := ctx.Query("author_id")
	str4 := ctx.Query("publisher_id")
	str5 := ctx.Query("genre_id")

	var authorId, publisherId, genreId int64
	if str3 == "" {
		authorId = 0
	} else {
		authorId, _ = strconv.ParseInt(str3, 10, 64)
	}

	if str4 == "" {
		publisherId = 0
	} else {
		publisherId, _ = strconv.ParseInt(str4, 10, 64)
	}

	if str5 == "" {
		genreId = 0
	} else {
		genreId, _ = strconv.ParseInt(str5, 10, 64)
	}

	books, dbError := handlers.GetBooks(searchValue, pageNumber, pageSize, authorId, publisherId, genreId)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"books": books})
}
