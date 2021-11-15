package api

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func createJWTToken(userName string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv("SPACELIGHT_JWT_SECRET"))
}
