package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var secret string = "9a0fd6dd6f38de8bd283c1f1aa5834d5a3d95f9839c6c7cc7aab138854fec50c768ad045c12e249adb3863d633b056ed27f20147b907019c0ee87d12ca6dcfd777048ccdbc43ee3b0ad6894a8d8be3ce0ad7da02bacc5093e5cc06f6b6cb694b4506ed5b26b5272b015c731f913fd25972734c3705ca756d3f804aa1abd0e5ac4f6b385113e72cc1539e0952f6119b1f591cc80c60bc27b2082196d473dfb09a1f7f01f642af498a4ce9bf110512610ee4919df46310b3016aa612aae50cc49ac87e150ea46b9628b88d3daba90e617972fda4e77710456d22461511716b714b4d08504442b3518d9a6fbb700d55e8c831b2d4bd2582dfd8b8f2a2a5e741d041"

func GenerateToken(username string) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	}).SignedString([]byte(secret))

	if err != nil {
		return "", errors.New(err.Error())
	}

	return token, nil
}
