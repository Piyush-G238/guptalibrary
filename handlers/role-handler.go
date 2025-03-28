package handlers

import (
	"errors"
	"time"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
)

func CreateRole(role *models.Role) (int, error) {

	fetchedRole := &models.Role{}
	configs.DB.Where("name = ?", role.Name).First(&fetchedRole)

	if fetchedRole.ID != 0 {
		return 0, errors.New("this role already exists")
	}
	role.CreatedAt = time.Now()
	configs.DB.Create(&role)
	return role.ID, nil
}
