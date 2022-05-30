package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var secret = os.Getenv("SPACELIGHT_JWT_SECRET")

const headerName = "auth"

type claims struct {
	UserId uuid.UUID `json:"UserId"`
	jwt.StandardClaims
}

func CreateJWTToken(userId uuid.UUID) (string, error) {
	log.Debugf("create JWT token")
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Errorf("Token creation failed, %v", err)
		return "", echo.ErrInternalServerError
	}

	log.Debugf("JWT token created")
	return signedToken, nil
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(headerName)
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		claims := &claims{}
		tkn, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			log.Debugf("Claim could not be parsed, %v", err)
			return echo.ErrUnauthorized
		}

		if tkn == nil || !tkn.Valid {
			log.Debugf("Token is not valid")
			return echo.ErrUnauthorized
		}

		if time.Now().Add(1 * time.Minute).After(time.Unix(claims.ExpiresAt, 0)) {
			token, err := CreateJWTToken(claims.UserId)
			if err != nil {
				return echo.ErrUnauthorized
			}
			c.Response().Header().Set(headerName, token)
		}
		return next(c)
	}
}
