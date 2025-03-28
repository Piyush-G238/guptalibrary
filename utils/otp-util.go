package utils

import (
	"crypto/rand"
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
