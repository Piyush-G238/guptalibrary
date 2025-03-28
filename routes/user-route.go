package routes

import (
	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/models"
)

func GroupUserRoutes(router *gin.RouterGroup) {
	router.POST("/signup", Signup)
	router.POST("/login", Login)
	router.POST("/verify-otp", VerifyOtp)
}

func Signup(ctx *gin.Context) {
	newUser := &models.User{}
	ctx.ShouldBindBodyWithJSON(newUser)

	newId, token, dbError := handlers.Signup(newUser)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": newId, "message": "User created successfully", "access_token": token})
}

func Login(ctx *gin.Context) {
	loginDetails := &models.User{}
	ctx.ShouldBindBodyWithJSON(loginDetails)

	newOtp, dbError := handlers.Login(loginDetails)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "User verified successfully!", "one_time_password": newOtp})
}

func VerifyOtp(ctx *gin.Context) {

	username := ctx.Query("username")
	otp := ctx.Query("otp")

	accessToken, dbError := handlers.VerifyOtp(username, otp)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OTP verified successfully!", "access_token": accessToken})
}
