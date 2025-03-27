package models

import "time"

type Author struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Books     []Book    `json:"books" foreign_key:"AuthorID"`
	CreatedAt time.Time `json:"created_at"`
}
