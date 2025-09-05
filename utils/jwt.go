package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("fsfdsfewerwtet57497yr")

type JWTClaims struct {
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(phone string, duration time.Duration) (string, error) {
	claims := JWTClaims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken() string {
	return fmt.Sprintf("%x", time.Now().UnixNano())
}

func ParseAccessToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
