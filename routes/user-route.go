package routes

import (
	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/models"
)

func GroupUserRoutes(router *gin.RouterGroup) {
	router.POST("/signup", Signup)
	router.POST("/login", Login)
	router.POST("/request-otp", RequestOtp)
	router.PATCH("/verify-login-otp", VerifyLoginOtp)
	router.PATCH("/verify-otp", VerifyOtp)
	router.PATCH("/verify-email", VerifyEmail)
	router.PATCH("/reset-password", ResetPassword)
}

func Signup(ctx *gin.Context) {
	newUser := &models.User{}
	ctx.ShouldBindBodyWithJSON(newUser)

	newId, dbError := handlers.Signup(newUser)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(201, gin.H{"id": newId, "message": "User created successfully and E-mail is sent for verification!"})
}

func Login(ctx *gin.Context) {
	loginDetails := &models.User{}
	ctx.ShouldBindBodyWithJSON(loginDetails)

	userName, dbError := handlers.Login(loginDetails)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "Credentials verified successfully!", "username": userName})
}

func VerifyLoginOtp(ctx *gin.Context) {

	username := ctx.Query("username")
	otp := ctx.Query("otp")

	accessToken, dbError := handlers.VerifyLoginOtp(username, otp)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OTP verified successfully!", "access_token": accessToken})
}

func VerifyOtp(ctx *gin.Context) {

	username := ctx.Query("username")
	otp := ctx.Query("otp")

	dbError := handlers.VerifyOtp(username, otp)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "OTP verified successfully!"})
}

func VerifyEmail(ctx *gin.Context) {

	verificationToken := ctx.Query("token")
	handlerError := handlers.VerifyEmail(verificationToken)

	if handlerError != nil {
		ctx.JSON(400, gin.H{"message": handlerError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "user's email is verified successfully!"})
}

func RequestOtp(ctx *gin.Context) {

	userName := ctx.Query("username")
	existingUser, dbError := handlers.RequestOtp(userName)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "OTP is sent successfully!", "username": existingUser})
}

func ResetPassword(ctx *gin.Context) {

	userName := ctx.Query("username")
	password := ctx.Query("password")

	dbError := handlers.ResetPassword(userName, password)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "password has been resetted successfully!"})
}
