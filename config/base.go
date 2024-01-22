package config

import "os"

const (
	SECRET_KEY = "9C4nkBP$7rzgPNzcyXHYL!CEFFH5b2UIZQw7US5vGcU!O%eWglqR5plwESKkuluA"
)

func GetSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey != "" {
		return secretKey
	}
	return SECRET_KEY
}
