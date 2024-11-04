package utils

import (
    "github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

func ValidateJWT(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
}
