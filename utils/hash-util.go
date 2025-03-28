package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	newPassword, hashingError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashingError != nil {
		return "", errors.New("Error hashing password")
	}
	return string(newPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
