package config

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var (
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

func GetSignatureKey() (res []byte) {
	key := os.Getenv("JWT_SIGNATURE_KEY")
	if key == "" {
		key = "secret"
	}
	res = []byte(key)
	return res
}
