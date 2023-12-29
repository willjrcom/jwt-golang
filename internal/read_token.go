package jwtconverter

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("chave_secreta_super_secreta")

func ReadToken(tokenString string) (jwt.MapClaims, error) {
	// Parse do token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Verificar se o token é válido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}
