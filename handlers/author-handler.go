package handlers

import (
	"errors"
	"strings"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func CreateAuthor(author *models.Author) (int, error) {

	configs.DB.Create(&author)
	return author.ID, nil
}

func UpdateAuthor(authorId int, author *models.Author) (int, error) {

	fetchedAuthor := &models.Author{}
	configs.DB.Where("id = ?", authorId).Find(&fetchedAuthor)

	if fetchedAuthor.ID == 0 {
		return 0, errors.New("author not found by ID, update failed")
	}

	configs.DB.Model(fetchedAuthor).Update("name", author.Name)
	return fetchedAuthor.ID, nil
}

func GetAuthors(searchValue string) ([]models.Author, error) {

	authors := []models.Author{}
	configs.DB.
		Where("lower(name) like ?", "%"+strings.ToLower(searchValue)+"%").
		Find(&authors)
	return authors, nil
}
