package utils

import (
	"go-template/pkg/envs"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    id,
		"user_email": email,
	})

	// Converte a chave para []byte
	secret := []byte(envs.GetEnvOrDie("JWT_SECRET"))

	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(envs.GetEnvOrDie("JWT_SECRET")), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
