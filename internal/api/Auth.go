package api

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type claims struct {
	UserId uuid.UUID `json:"UserId"`
	jwt.StandardClaims
}

func createJWTToken(userId uuid.UUID) (string, error) {
	log.Debugf("create JWT token")
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SPACELIGHT_JWT_SECRET")))
	if err != nil {
		log.Errorf("Token creation failed, %v", err)
		return "", echo.ErrInternalServerError
	}

	log.Debugf("JWT token created")
	return signedToken, nil
}
