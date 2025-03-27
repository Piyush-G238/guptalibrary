package models

import "time"

type Role struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Users     []User    `json:"users" gorm:"many2many:user_roles;"`
}
