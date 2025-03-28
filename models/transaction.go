package models

import "time"

type Transaction struct {
	ID         int       `json:"id" gorm:"primary_key"`
	UserId     int       `json:"user_id"`
	BookId     int       `json:"book_id"`
	IssueDate  time.Time `json:"issue_date"`
	ReturnDate time.Time `json:"return_date"`
	Status     string    `json:"status"`
	User       User      `json:"user"`
	Book       Book      `json:"book"`
}
