package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToke(jwtKey string, issuer string) (string, error) {

	claims := jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(jwtKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
