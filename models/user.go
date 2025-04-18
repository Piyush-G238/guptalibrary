package models

import "time"

type User struct {
	ID              int       `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	Password        string    `json:"password"`
	CreatedAt       time.Time `json:"created_at"`
	Roles           []Role    `json:"roles" gorm:"many2many:user_roles;"`
	IsEmailVerified bool      `json:"is_email_verified"`
}
