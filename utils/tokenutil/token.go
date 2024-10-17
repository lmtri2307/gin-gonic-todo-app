package tokenutil

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(payload any, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := jwt.MapClaims{
		"payload": payload,
		"exp":     exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
