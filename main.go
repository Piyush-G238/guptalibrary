package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/configs"
	"guptalibrary.com/routes"
)

func main() {

	fmt.Println("Hello, World!")

	configs.InitDB()
	configs.InitRedis()

	application := gin.Default()

	GroupAllRoutes(application.Group("/api/v1"))
	application.Run(":8080")
}

func GroupAllRoutes(router *gin.RouterGroup) {
	routes.GroupPublisherRoutes(router.Group("/publishers"))
	routes.GroupAuthorRoutes(router.Group("/authors"))
	routes.GroupBookRoutes(router.Group("/books"))
	routes.GroupGenreRoutes(router.Group("/genres"))
	routes.GroupRoleRoute(router.Group("/roles"))
	routes.GroupUserRoutes(router.Group("/users"))
	routes.GroupTransactionRoutes(router.Group("/transactions"))
}
