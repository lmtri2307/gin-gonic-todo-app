package tokenutil

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(payload any, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Second * time.Duration(expiry)).Unix()
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

func VerifyToken(tokenString string, secret string) (any, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		hmacSampleSecret := []byte(secret)
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["payload"], nil
	} else {
		return nil, err
	}
}
