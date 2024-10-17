package my_jwt

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"user_name"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}
