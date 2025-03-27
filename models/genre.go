package models

import "time"

type Genre struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Books     []Book    `json:"books" gorm:"many2many:book_genres"`
}
