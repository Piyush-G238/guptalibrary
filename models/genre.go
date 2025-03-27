package models

type Genre struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Books []Book `json:"books" gorm:"many2many:book_genres"`
}
