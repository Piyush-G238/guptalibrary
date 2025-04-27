package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
	"guptalibrary.com/middlewares"
)

func GroupTransactionRoutes(router *gin.RouterGroup) {

	router.
		Use(middlewares.AuthenticationMiddleware()).
		Use(middlewares.AdminMiddleware()).
		POST("/admin/issue-book", IssueBookByAdmin)
	router.
		Use(middlewares.AuthenticationMiddleware()).
		Use(middlewares.AdminMiddleware()).
		PATCH("/admin/return-book", ReturnBookByAdmin)
	router.
		Use(middlewares.AuthenticationMiddleware()).
		Use(middlewares.AdminMiddleware()).
		PATCH("/admin/cancel-reservation/:id", CancelReservation)
	router.
		Use(middlewares.AuthenticationMiddleware()).
		Use(middlewares.AdminMiddleware()).
		PATCH("/admin/notify-reservation/:id", NotifyReservation)
}

func IssueBookByAdmin(ctx *gin.Context) {

	bookIdParam := ctx.Query("book_id")
	issueDate := ctx.Query("issue_date")
	returnDate := ctx.Query("due_date")
	userIdParam := ctx.Query("user_id")

	userId, parsedErr := strconv.ParseInt(userIdParam, 10, 64)
	if parsedErr != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user id"})
		return
	}

	bookId, parsedErr := strconv.ParseInt(bookIdParam, 10, 64)
	if parsedErr != nil {
		ctx.JSON(400, gin.H{"error": "Invalid book id"})
		return
	}

	transactionId, dbError := handlers.IssueBookByAdmin(int(userId), int(bookId), issueDate, returnDate)
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"transaction_id": transactionId, "message": "Book issued successfully"})
}

func ReturnBookByAdmin(ctx *gin.Context) {

	transactionIdParam := ctx.Query("transaction_id")

	transactionId, parsedErr := strconv.ParseInt(transactionIdParam, 10, 64)
	if parsedErr != nil {
		ctx.JSON(400, gin.H{"error": "Invalid transaction id"})
		return
	}

	txnId, dbError := handlers.ReturnBookByAdmin(int(transactionId))

	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"transaction_id": txnId, "message": "Book returned successfully"})
}

func CancelReservation(ctx *gin.Context) {

	str1 := ctx.Param("id")
	reservationId, parseError := strconv.ParseInt(str1, 10, 64)

	if parseError != nil {
		ctx.JSON(400, gin.H{"error": parseError.Error()})
	}

	_, dbError := handlers.CancelReservation(int(reservationId))
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "reservation has been cancelled successfully"})
}

func NotifyReservation(ctx *gin.Context) {

	str1 := ctx.Param("id")
	reservationId, parseError := strconv.ParseInt(str1, 10, 64)

	if parseError != nil {
		ctx.JSON(400, gin.H{"error": parseError.Error()})
	}

	_, dbError := handlers.NotifyReservation(int(reservationId))
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Notification has been sent successfully for this reservation"})
}
