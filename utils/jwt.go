package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "p5X957pAPBU+ixjePRYXCOM+ZRaWuPMtPuKmNn24ohM="

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC) // type checking

			if !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(secretKey), nil
		})

	if err != nil {
		return 0, errors.New("could not parse provided token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) // type checking

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string) // type checking
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
