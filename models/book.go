package models

import "time"

type Book struct {
	ID              int       `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	AuthorID        int       `json:"author_id"`
	Isbn            string    `json:"isbn"`
	PublishedYear   int       `json:"published_year"`
	TotalCopies     int       `json:"total_copies"`
	AvailableCopies int       `json:"available_copies"`
	Price           float64   `json:"price"`
	Author          Author    `json:"author"`
	PublisherID     int       `json:"publisher_id"`
	Publisher       Publisher `json:"publisher"`
	Genres          []Genre   `gorm:"many2many:book_genres;" json:"genres"`
	CreatedAt       time.Time `json:"created_at"`
}
