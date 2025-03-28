package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"guptalibrary.com/handlers"
)

func GroupTransactionRoutes(router *gin.RouterGroup) {

	router.POST("/admin/issue-book", IssueBookByAdmin)    // for admin, to issue a book on the behalf of user
	router.PATCH("/admin/return-book", ReturnBookByAdmin) // for admin, to return a book on the behalf of user
	router.GET("/", GetTransactions)                      // get all transactions
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
	returnDate := ctx.Query("return_date")

	transactionId, parsedErr := strconv.ParseInt(transactionIdParam, 10, 64)
	if parsedErr != nil {
		ctx.JSON(400, gin.H{"error": "Invalid transaction id"})
		return
	}

	txnId, dbError := handlers.ReturnBookByAdmin(int(transactionId), returnDate)

	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"transaction_id": txnId, "message": "Book returned successfully"})
}

func GetTransactions(ctx *gin.Context) {

	transactions, dbError := handlers.GetTransactions()
	if dbError != nil {
		ctx.JSON(400, gin.H{"error": dbError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"transactions": transactions})
}
