package utils

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strconv"
)

func GenerateOTP() string {
	otp := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		otp += strconv.Itoa(int(n.Int64()))
	}
	return otp
}

func GenerateVerificationToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
