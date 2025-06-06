package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/middlewares"
	"guptalibrary.com/models"
)

func GroupPublisherRoutes(router *gin.RouterGroup) {

	router.
		Use(middlewares.AuthenticationMiddleware()).
		Use(middlewares.AdminMiddleware()).
		POST("/", CreatePublisher)
	router.
		Use(middlewares.AuthenticationMiddleware()).
		Use(middlewares.AdminMiddleware()).
		PATCH("/:id", UpdatePublisher)
	router.
		Use(middlewares.AuthenticationMiddleware()).
		GET("/", GetPublishers)
}

func CreatePublisher(ctx *gin.Context) {
	publisher := &models.Publisher{}
	ctx.ShouldBindBodyWithJSON(publisher)

	generatedId, dbError := handlers.CreatePublisher(publisher)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}
	ctx.JSON(201, gin.H{"id": generatedId, "message": "Publisher created successfully!"})
}

func UpdatePublisher(ctx *gin.Context) {
	value, doesExits := ctx.Params.Get("id")
	if !doesExits {
		ctx.JSON(400, gin.H{"error": "Param Id not available"})
		return
	}
	publisherId, parseError := strconv.ParseInt(value, 10, 64)
	if parseError != nil {
		ctx.JSON(400, gin.H{"error": "Unable to parse Id"})
		return
	}

	publisher := &models.Publisher{}
	ctx.ShouldBindBodyWithJSON(publisher)

	generatedId, dbError := handlers.UpdatePublisher(int(publisherId), publisher)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": generatedId, "message": "Publisher updated successfully!"})
}

func GetPublishers(ctx *gin.Context) {
	searchValue := ctx.Query("search")

	publishers := handlers.GetPublishers(searchValue)
	ctx.JSON(200, gin.H{"publishers": publishers})
}
