package models

type Publisher struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Books []Book `json:"books" gorm:"foreignKey:PublisherID"`
}
