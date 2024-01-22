package lib

import "github.com/golang-jwt/jwt"

type MyClaim struct {
	jwt.StandardClaims
	Roles  []int  `json:"roles"`
	Email  string `json:"email"`
	UserId uint   `json:"user_id"`
}
