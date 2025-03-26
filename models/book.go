package models

type Book struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	AuthorID    int       `json:"author_id"`
	Author      Author    `json:"author"`
	PublisherID int       `json:"publisher_id"`
	Publisher   Publisher `json:"publisher"`
	// GenreIDs    []int     `json:"genre_ids"`
	Genres []Genre `gorm:"many2many:book_genres;" json:"genres"`
}
