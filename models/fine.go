package models

import "time"

type Fine struct {
	ID            int         `json:"id"`
	UserId        int         `json:"user_id"`
	TransactionId int         `json:"transaction_id"`
	Amount        float64     `json:"amount"`
	Paid          bool        `json:"paid"`
	ImposedDate   time.Time   `json:"imposed_date"`
	PaidDate      time.Time   `json:"paid_date"`
	User          User        `json:"user"`
	Transaction   Transaction `json:"transaction"`
	CreatedAt     time.Time   `json:"created_at"`
}
