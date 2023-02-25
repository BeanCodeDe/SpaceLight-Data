package util

import (
	"crypto/rsa"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	PublicKeyFile  = "./../deployments/token/jwtRS256.key.pub"
	PrivateKeyFile = "./../deployments/token/jwtRS256.key"
)

type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func CreateUserId() string {
	return uuid.New().String()
}

func CreateJWTToken(userId string) string {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(loadPrivateKeyFile(PrivateKeyFile))
	if err != nil {
		panic(err)
	}

	return signedToken
}

func loadPrivateKeyFile(fileName string) *rsa.PrivateKey {
	verifyBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	verifyKey, err := jwt.ParseRSAPrivateKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}

	return verifyKey
}
