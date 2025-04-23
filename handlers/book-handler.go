package handlers

import (
	"errors"
	"strings"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func CreateBook(newBook *models.Book) (int, error) {

	fetchedBook := &models.Book{}
	configs.DB.Where("lower(name) = ? and author_id = ?", strings.ToLower(newBook.Name), newBook.AuthorID).First(fetchedBook)

	if fetchedBook.ID != 0 {
		return 0, errors.New("book is already present with the same name and author")
	}

	author := &models.Author{}
	configs.DB.Where("id = ?", newBook.AuthorID).Select("id").First(&author)

	if author.ID == 0 {
		return 0, errors.New("author is not present")
	}

	publisher := &models.Publisher{}
	configs.DB.Where("id = ?", newBook.PublisherID).Select("id").First(&publisher)

	if publisher.ID == 0 {
		return 0, errors.New("publisher is not present")
	}

	configs.DB.Create(newBook)
	return newBook.ID, nil
}

func UpdateBook(bookId int, book *models.Book) (int, error) {

	fetchedBook := &models.Book{}
	configs.DB.Where("id = ?", bookId).First(fetchedBook)
	if fetchedBook.ID == 0 {
		return 0, errors.New("book is not present")
	}

	existingBook := &models.Book{}
	configs.DB.Where("lower(name) = ? and author_id = ? and id != ?", strings.ToLower(book.Name), book.AuthorID, bookId).First(existingBook)
	if existingBook.ID != 0 {
		return 0, errors.New("book is already present with the same name and author")
	}

	book.ID = fetchedBook.ID
	configs.DB.Save(book)
	return fetchedBook.ID, nil
}

func GetBooks(searchValue string, pageNumber, pageSize, authorId, publisherId, genreId int64) ([]models.Book, error) {

	books := []models.Book{}

	txn := configs.DB.
		Where(
			`(lower(books.name) like ? or lower(books.isbn) like ?)`,
			strings.ToLower("%"+searchValue+"%"),
			strings.ToLower("%"+searchValue+"%"))

	if authorId != 0 {
		txn.Where("books.author_id = ?", authorId)
	}

	if publisherId != 0 {
		txn.Where("books.publisher_id = ?", publisherId)
	}

	if genreId != 0 {
		txn.Where("book_genres.genre_id = ?", genreId)
	}

	txn.
		Joins("JOIN book_genres on book_genres.book_id = books.id").
		Limit(int(pageSize)).
		Offset(int(pageNumber-1) * int(pageSize)).
		Group("books.id").
		Preload("Author").
		Find(&books)

	return books, nil
}
