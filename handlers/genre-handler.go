package handlers

import (
	"errors"
	"strings"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func CreateGenre(genre *models.Genre) (int, error) {

	fetchedGenre := &models.Genre{}
	configs.DB.Where("lower(name) = ?", strings.ToLower(genre.Name)).Find(fetchedGenre)

	if fetchedGenre.ID != 0 {
		return 0, errors.New("genre is already present with the same name")
	}

	configs.DB.Create(genre)
	return genre.ID, nil
}

func GetGenres() ([]models.Genre, error) {

	var genres []models.Genre
	configs.DB.Preload("Books", []models.Book{}).Find(&genres)
	return genres, nil
}
