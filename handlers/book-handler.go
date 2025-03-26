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
