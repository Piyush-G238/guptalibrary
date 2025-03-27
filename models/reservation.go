package models

import "time"

type Reservation struct {
	ID              int       `json:"id"`
	UserId          int       `json:"user_id"`
	BookId          int       `json:"book_id"`
	ReservationDate time.Time `json:"reservation_date"`
	Status          string    `json:"status"`
	User            User      `json:"user" gorm:"foreignKey:ID"`
	Book            Book      `json:"book" gorm:"foreignKey:ID"`
	CreatedAt       time.Time `json:"created_at"`
}
