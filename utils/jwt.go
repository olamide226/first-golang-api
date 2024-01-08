package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

type ParsedClaims struct {
	Email  string
	UserID int64
}

func GenerateToken(id int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (*ParsedClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("Could not parse token")
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email := claims["email"].(string)
		userId := claims["userId"].(float64)
		return &ParsedClaims{
			Email:  email,
			UserID: int64(userId),
		}, nil
	}
	return nil, errors.New("Could not parse claims")
}

func SplitToken(token string) string {
	if len(token) < 7 {
		return ""
	}
	return token[7:]
}
