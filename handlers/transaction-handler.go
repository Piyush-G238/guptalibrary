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
		return 0, errors.New("user not found")
	}

	if fetchedBook.ID == 0 {
		return 0, errors.New("book not found")
	}

	if fetchedBook.AvailableCopies == 0 {
		// create a new reservation for this user
		newReservation := models.Reservation{}
		newReservation.UserId = userId
		newReservation.BookId = bookId
		newReservation.Status = "pending"
		newReservation.ReservationDate = time.Now()
		configs.DB.Create(&newReservation)

		dynamicValues := make(map[string]any)
		dynamicValues["Username"] = fetchedUser.Username
		dynamicValues["BookName"] = fetchedBook.Name
		dynamicValues["IsbnNumber"] = fetchedBook.Isbn

		reservationDate := newReservation.ReservationDate.Format("2006-01-02")
		dynamicValues["ReservationDate"] = reservationDate

		_, emailError := SendEmail("reservation confirmed template", "Reservation confirmed",
			dynamicValues, fetchedUser.Email)

		if emailError != nil {
			return 0, errors.New("Unable to send error: " + emailError.Error())
		}

		return 0, errors.New("book is not available to issue")
	}

	transaction := models.Transaction{}
	transaction.UserId = userId
	transaction.BookId = bookId

	issueDateParsed, err := time.Parse("2006-01-02", issueDate)
	if err != nil {
		return 0, errors.New("invalid issue date")
	}
	transaction.IssueDate = issueDateParsed

	dueDateParsed, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return 0, errors.New("invalid return date")
	}
	transaction.DueDate = dueDateParsed
	transaction.Status = "borrowed"
	transaction.ReturnDate = time.Time{}

	configs.DB.Model(&fetchedBook).Update("available_copies", fetchedBook.AvailableCopies-1)
	configs.DB.Create(&transaction)

	dynamicValues := make(map[string]any)
	dynamicValues["Username"] = fetchedUser.Username
	dynamicValues["BookName"] = fetchedBook.Name
	dynamicValues["IsbnNumber"] = fetchedBook.Isbn
	dynamicValues["IssueDate"] = issueDate
	dynamicValues["DueDate"] = dueDate

	_, emailError := SendEmail("issued book template", "Issued book successfully",
		dynamicValues, fetchedUser.Email)

	if emailError != nil {
		configs.DB.Rollback()
		return 0, errors.New("Unable to send email, error: " + emailError.Error())
	}
	return transaction.ID, nil
}

func ReturnBookByAdmin(transactionId int) (int, error) {
	fetchedTransaction := models.Transaction{}
	configs.DB.
		Where("id = ?", transactionId).
		Preload("User").
		Preload("Book").
		First(&fetchedTransaction)

	if fetchedTransaction.ID == 0 {
		return 0, errors.New("transaction not found")
	}

	if fetchedTransaction.Status == "returned" {
		return 0, errors.New("book is already returned")
	}

	configs.DB.Model(&fetchedTransaction).Update("return_date", time.Now())
	configs.DB.Model(&fetchedTransaction).Update("status", "returned")

	fetchedBook := models.Book{}
	configs.DB.Where("id = ?", fetchedTransaction.BookId).First(&fetchedBook)
	configs.DB.Model(&fetchedBook).Update("available_copies", fetchedBook.AvailableCopies+1)

	// notify user for book availability
	latestReservation := &models.Reservation{}
	configs.DB.
		Where("book_id = ? and status = ?", fetchedTransaction.BookId, "pending").
		Order("reservation_date desc").
		Preload("User").
		Find(latestReservation)

	if latestReservation.ID != 0 {
		latestReservation.Status = "notified"
		configs.DB.Save(latestReservation)

		dynamicValues := make(map[string]any)
		dynamicValues["Username"] = latestReservation.User.Username
		dynamicValues["BookName"] = fetchedBook.Name

		_, emailError := SendEmail("book available notify template", "Book Available to Issue",
			dynamicValues, fetchedTransaction.User.Email)

		if emailError != nil {
			configs.DB.Rollback()
			return 0, errors.New("Unable to send email, error: " + emailError.Error())
		}
	}

	dynamicValues := make(map[string]any)
	dynamicValues["Username"] = fetchedTransaction.User.Username
	dynamicValues["BookName"] = fetchedBook.Name
	dynamicValues["IsbnNumber"] = fetchedBook.Isbn
	dynamicValues["IssueDate"] = fetchedTransaction.IssueDate.Format("2006-01-01")
	dynamicValues["ReturnDate"] = fetchedTransaction.ReturnDate.Format("2006-01-01")

	_, emailError := SendEmail("book returned template", "Book Returned successfully",
		dynamicValues, fetchedTransaction.User.Email)

	if emailError != nil {
		configs.DB.Rollback()
		return 0, errors.New("Unable to send email, error: " + emailError.Error())
	}

	return transactionId, nil

}

func CancelReservation(reservationId int) (int, error) {

	fetchedReservation := &models.Reservation{}
	configs.DB.
		Where("id = ?", reservationId).
		First(fetchedReservation)

	if fetchedReservation.ID == 0 {
		return 0, errors.New("reservation not found with provided ID")
	}

	if fetchedReservation.Status == "cancelled" {
		return 0, errors.New("this reservation is already cancelled")
	}

	fetchedReservation.Status = "cancelled"
	configs.DB.Save(fetchedReservation)
	return reservationId, nil
}

func NotifyReservation(reservationId int) (int, error) {

	fetchedReservation := &models.Reservation{}
	configs.DB.
		Where("id = ?", reservationId).
		Preload("User").
		Preload("Book").
		First(fetchedReservation)

	if fetchedReservation.ID == 0 {
		return 0, errors.New("reservation not found with provided ID")
	}

	if fetchedReservation.Status == "cancelled" {
		return 0, errors.New("this reservation is already cancelled")
	} else if fetchedReservation.Status == "notified" {
		return 0, errors.New("notification is already sent for this reservation")
	}

	fetchedReservation.Status = "notified"
	configs.DB.Save(fetchedReservation)

	dynamicValues := make(map[string]any)
	dynamicValues["Username"] = fetchedReservation.User.Username
	dynamicValues["BookName"] = fetchedReservation.Book.Name

	_, emailError := SendEmail("book available notify template", "Book Available to Issue",
		dynamicValues, fetchedReservation.User.Email)

	if emailError != nil {
		configs.DB.Rollback()
		return 0, errors.New("Unable to send email, error: " + emailError.Error())
	}
	return reservationId, nil
}
