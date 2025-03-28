package handlers

import (
	"errors"
	"time"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func IssueBookByAdmin(userId, bookId int, issueDate, dueDate string) (int, error) {

	fetchedUser := models.User{}
	fetchedBook := models.Book{}
	configs.DB.Where("id = ?", userId).First(&fetchedUser)
	configs.DB.Where("id = ?", bookId).First(&fetchedBook)

	if fetchedUser.ID == 0 {
		return 0, errors.New("User not found")
	}

	if fetchedBook.ID == 0 {
		return 0, errors.New("Book not found")
	}

	if fetchedBook.AvailableCopies == 0 {
		// create a new reservation for this user
		newReservation := models.Reservation{}
		newReservation.UserId = userId
		newReservation.BookId = bookId
		newReservation.Status = "pending"
		newReservation.ReservationDate = time.Now()
		configs.DB.Create(&newReservation)

		return 0, errors.New("Book is not available to issue")
	}

	transaction := models.Transaction{}
	transaction.UserId = userId
	transaction.BookId = bookId

	issueDateParsed, err := time.Parse("2006-01-02", issueDate)
	if err != nil {
		return 0, errors.New("Invalid issue date")
	}
	transaction.IssueDate = issueDateParsed

	dueDateParsed, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return 0, errors.New("Invalid return date")
	}
	transaction.DueDate = dueDateParsed
	transaction.Status = "borrowed"
	transaction.ReturnDate = time.Time{}

	configs.DB.Model(&fetchedBook).Update("available_copies", fetchedBook.AvailableCopies-1)
	configs.DB.Create(&transaction)
	return transaction.ID, nil
}

func ReturnBookByAdmin(transactionId int, returnDate string) (int, error) {
	fetchedTransaction := models.Transaction{}
	configs.DB.Where("id = ?", transactionId).First(&fetchedTransaction)

	if fetchedTransaction.ID == 0 {
		return 0, errors.New("Transaction not found")
	}

	if fetchedTransaction.Status == "returned" {
		return 0, errors.New("Book is already returned")
	}

	returnDateParsed, err := time.Parse("2006-01-02", returnDate)
	if err != nil {
		return 0, errors.New("Invalid return date")
	}

	configs.DB.Model(&fetchedTransaction).Update("return_date", returnDateParsed)
	configs.DB.Model(&fetchedTransaction).Update("status", "returned")

	fetchedBook := models.Book{}
	configs.DB.Where("id = ?", fetchedTransaction.BookId).First(&fetchedBook)
	configs.DB.Model(&fetchedBook).Update("available_copies", fetchedBook.AvailableCopies+1)

	return transactionId, nil

}

func GetTransactions() ([]models.Transaction, error) {
	transactions := []models.Transaction{}
	configs.DB.Preload("User", &models.User{}).Preload("Book", &models.Book{}).Find(&transactions)
	return transactions, nil
}
