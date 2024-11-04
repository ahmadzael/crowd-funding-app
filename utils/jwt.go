package utils

import (
    "github.com/golang-jwt/jwt"
    "time"
)

var jwtKey = []byte("my_secret_key")

func GenerateJWT(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString(jwtKey)
}
