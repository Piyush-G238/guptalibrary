package handlers

import (
	"errors"
	"time"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
	"guptalibrary.com/utils"
)

func Signup(newUser *models.User) (int, string, error) {

	userExists := &models.User{}
	configs.DB.Where("username = ?", newUser.Username).First(userExists)

	if userExists.ID != 0 {
		return 0, "", errors.New("Username already exists")
	}
	configs.DB.Where("email = ?", newUser.Email).First(userExists)
	if userExists.ID != 0 {
		return 0, "", errors.New("Email already exists")
	}
	configs.DB.Where("phone = ?", newUser.Phone).First(userExists)
	if userExists.ID != 0 {
		return 0, "", errors.New("Phone already exists")
	}

	hashedPassword, hashingError := utils.HashPassword(newUser.Password)
	if hashingError != nil {
		return 0, "", errors.New("Error hashing password")
	}
	newUser.Password = hashedPassword

	newUser.Roles = make([]models.Role, 1)
	configs.DB.Where("name = ?", "MEMBER").First(&newUser.Roles[0])

	newUser.CreatedAt = configs.DB.NowFunc()
	configs.DB.Create(newUser)

	token, tokenError := utils.GenerateToken(newUser.Username)
	if tokenError != nil {
		return 0, "", errors.New("Error generating token")
	}
	return newUser.ID, token, nil
}

func Login(user *models.User) (string, error) {

	userExists := &models.User{}
	configs.DB.Where("username = ?", user.Username).First(userExists)

	if userExists.ID == 0 {
		return "", errors.New("User not found")
	}

	passwordMatch := utils.CheckPasswordHash(user.Password, userExists.Password)
	if !passwordMatch {
		return "", errors.New("Invalid password")
	}

	newOtp := utils.GenerateOTP()
	configs.RedisClient.Set(configs.Context, user.Username, newOtp, 5*time.Minute)
	return newOtp, nil
}

func VerifyOtp(username, otp string) (string, error) {

	storedOtp, err := configs.RedisClient.Get(configs.Context, username).Result()
	if err != nil {
		return "", errors.New("OTP expired/OTP not generated for the user")
	}

	if storedOtp != otp {
		return "", errors.New("Invalid OTP")
	}

	configs.RedisClient.Del(configs.Context, username)

	token, tokenError := utils.GenerateToken(username)
	if tokenError != nil {
		return "", errors.New("Error generating token")
	}
	return token, nil
}
