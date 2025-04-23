package handlers

import (
	"fmt"
	"strings"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func CreatePublisher(publisher *models.Publisher) (int, error) {

	fetchedPublisher := &models.Publisher{}
	configs.DB.Where("lower(name) = ?", strings.ToLower(publisher.Name)).Select("id").Find(fetchedPublisher)
	if fetchedPublisher.ID != 0 {
		return 0, fmt.Errorf("publisher with name %s already exists", publisher.Name)
	}
	configs.DB.Create(&publisher)
	return publisher.ID, nil
}

func UpdatePublisher(publisherId int, publisher *models.Publisher) (int, error) {

	fetchedPublisher := &models.Publisher{}
	configs.DB.Where("id = ?", publisherId).First(fetchedPublisher)

	if fetchedPublisher.ID == 0 {
		return 0, fmt.Errorf("publisher with id %d not found", publisherId)
	}
	anyOtherPublisher := &models.Publisher{}
	configs.DB.Where("lower(name) = ? AND id != ?", strings.ToLower(publisher.Name), publisherId).Select("id").Find(anyOtherPublisher)
	if anyOtherPublisher.ID != 0 {
		return 0, fmt.Errorf("publisher with name '%s' already exists", publisher.Name)
	}
	configs.DB.Model(&fetchedPublisher).Update("name", publisher.Name)
	return fetchedPublisher.ID, nil
}

func GetPublishers() []models.Publisher {
	publishers := []models.Publisher{}
	configs.DB.Find(&publishers)
	return publishers
}
