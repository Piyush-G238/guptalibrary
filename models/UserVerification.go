package models

import "time"

type UserVerification struct {
	UserId            int       `json:"user_id"`
	VerificationToken string    `json:"verification_token"`
	ExpirationTime    time.Time `json:"expiration_time"`
}
