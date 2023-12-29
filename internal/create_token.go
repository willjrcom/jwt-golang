package jwtconverter

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Função para criar um token
func CreateToken() (string, error) {
	// Criar uma estrutura de claims
	claims := jwt.MapClaims{
		"name":  "will",
		"age":   23,
		"email": "will@will.com",
		"exp":   time.Now().Add(time.Second * 24).Unix(), // Token expira em 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
