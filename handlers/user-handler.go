package handlers

import (
	"errors"
	"strings"
	"time"

	"guptalibrary.com/configs"
	"guptalibrary.com/models"
	"guptalibrary.com/utils"
)

func Signup(newUser *models.User) (int, error) {

	userExists := &models.User{}
	configs.DB.Where("username = ?", newUser.Username).First(userExists)

	if userExists.ID != 0 {
		return 0, errors.New("username already exists")
	}
	configs.DB.Where("email = ?", newUser.Email).First(userExists)

	if userExists.ID != 0 {
		return 0, errors.New("email already exists")
	}

	if strings.Trim(newUser.Phone, " ") != "" {
		configs.DB.Where("phone = ?", newUser.Phone).First(userExists)
		if userExists.ID != 0 {
			return 0, errors.New("phone already exists")
		}
	}

	hashedPassword, hashingError := utils.HashPassword(newUser.Password)
	if hashingError != nil {
		return 0, errors.New("error while hashing password")
	}
	newUser.Password = hashedPassword

	newUser.Roles = make([]models.Role, 1)
	configs.DB.Where("name = ?", "MEMBER").First(&newUser.Roles[0])

	newUser.CreatedAt = configs.DB.NowFunc()
	configs.DB.Create(newUser)

	verifyToken, tokenError := utils.GenerateVerificationToken()
	// token, tokenError := utils.GenerateToken(newUser.Username)
	if tokenError != nil {
		return 0, errors.New("error while generating verification token")
	}

	userVerification := &models.UserVerification{
		UserId:            newUser.ID,
		VerificationToken: verifyToken,
		ExpirationTime:    time.Now().Add(time.Hour * 24),
	}

	configs.DB.Create(userVerification)
	dynamicValues := make(map[string]any)
	dynamicValues["Username"] = newUser.Username
	dynamicValues["VerificationLink"] = "http://localhost:8080/api/v1/users/verify-email?token=" + verifyToken

	_, emailError := SendEmail(
		"verify email address template",
		"Verify Email Address",
		dynamicValues,
		newUser.Email)

	if emailError != nil {
		return 0, errors.New("Unable to send error: " + emailError.Error())
	}

	return newUser.ID, nil
}

func Login(user *models.User) (string, error) {

	userExists := &models.User{}
	configs.DB.Where("username = ? or email = ?", user.Username, user.Username).First(userExists)

	if userExists.ID == 0 {
		return "", errors.New("invalid credentials")
	}

	if !userExists.IsEmailVerified {
		return "", errors.New("email is not verified yet, please verify your email first")
	}

	passwordMatch := utils.CheckPasswordHash(user.Password, userExists.Password)
	if !passwordMatch {
		return "", errors.New("invalid credentials")
	}

	newOtp := utils.GenerateOTP()
	configs.RedisClient.Set(configs.Context, "otp_"+userExists.Username, newOtp, 5*time.Minute)

	dynamicValues := make(map[string]any)
	dynamicValues["Username"] = userExists.Username
	dynamicValues["OTP"] = newOtp

	_, emailError := SendEmail(
		"login otp template",
		"Verify OTP",
		dynamicValues,
		userExists.Email)

	if emailError != nil {
		return "", errors.New("Unable to send error: " + emailError.Error())
	}

	return userExists.Username, nil
}

func VerifyLoginOtp(username, otp string) (string, error) {

	otpError := VerifyOtp(username, otp)
	if otpError != nil {
		return "", otpError
	}

	token, tokenError := utils.GenerateToken(username)
	if tokenError != nil {
		return "", errors.New("error generating token")
	}
	return token, nil
}

func RequestOtp(username string) (string, error) {

	userExists := &models.User{}
	configs.DB.Where("username = ? or email = ?", username, username).First(userExists)

	if userExists.ID == 0 {
		return "", errors.New("username/email not found")
	}

	if !userExists.IsEmailVerified {
		return "", errors.New("email is not verified yet, please verify your email first")
	}

	newOtp := utils.GenerateOTP()
	configs.RedisClient.Set(configs.Context, "otp_"+userExists.Username, newOtp, 5*time.Minute)

	dynamicValues := make(map[string]any)
	dynamicValues["Username"] = userExists.Username
	dynamicValues["OTP"] = newOtp

	_, emailError := SendEmail(
		"login otp template",
		"Verify OTP",
		dynamicValues,
		userExists.Email)

	if emailError != nil {
		return "", errors.New("Unable to send error: " + emailError.Error())
	}

	return userExists.Username, nil
}

func VerifyOtp(username, otp string) error {

	storedOtp, err := configs.RedisClient.Get(configs.Context, "otp_"+username).Result()
	if err != nil {
		return errors.New("OTP expired/OTP not generated for the user")
	}

	if storedOtp != otp {
		return errors.New("invalid OTP provided")
	}

	configs.RedisClient.Del(configs.Context, "otp_"+username)
	return nil
}

func VerifyEmail(token string) error {

	verification := &models.UserVerification{}
	configs.DB.Where("verification_token = ?", token).Find(verification)

	if verification.UserId == 0 {
		return errors.New("no user is found against this verification token")
	}

	expiryTime := verification.ExpirationTime
	if expiryTime.Before(time.Now()) {
		configs.DB.Delete(verification)
		return errors.New("verification token is expired, please request for email verification again")
	} else if verification.VerificationToken != token {
		return errors.New("wrong verification token, please request for email verification again")
	}

	updatedUser := &models.User{
		ID:              verification.UserId,
		IsEmailVerified: true,
	}
	// configs.DB.Model(updatedUser).Where("id = ?", updatedUser.ID).UpdateColumn("is_email_verified = ?", updatedUser.IsEmailVerified)
	configs.DB.Model(updatedUser).Where("id = ?", updatedUser.ID).UpdateColumn("is_email_verified", updatedUser.IsEmailVerified)
	configs.DB.Where("user_id = ?", verification.UserId).Delete(verification)

	return nil
}

func ResetPassword(username, password string) error {

	userExists := &models.User{}
	configs.DB.Where("username = ?", username).First(userExists)

	if userExists.ID == 0 {
		return errors.New("username/email not found")
	}

	encryptedPassword, _ := utils.HashPassword(password)
	userExists.Password = encryptedPassword

	configs.DB.Model(userExists).Where("id = ?", userExists.ID).UpdateColumn("password", userExists.Password)
	return nil
}
