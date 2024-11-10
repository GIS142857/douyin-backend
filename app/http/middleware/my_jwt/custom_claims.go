package my_jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UID      int64  `json:"uid"`
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
	Password string `json:"password"` // varchar(100)
	jwt.StandardClaims
}
