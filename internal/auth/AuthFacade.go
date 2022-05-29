package auth

import (
	"net/http"
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
		if time.Now().Before(time.Unix(claims.ExpiresAt, 0)) {
			log.Debugf("Token is passed expire date")
			return echo.ErrUnauthorized
		}
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

		token, err := CreateJWTToken(claims.UserId)
		if err != nil {
			return echo.ErrUnauthorized
		}
		c.Response().Header().Set(headerName, token)

		return next(c)
	}
}

func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If the user is not authenticated (no user token data in the context), don't do anything.
		authHeader := c.Request().Header.Get(headerName)
		if authHeader == "" {
			return next(c)
		}

		claims := claims{}
		// We ensure that a new token is not issued until enough time has elapsed.
		// In this case, a new token will only be issued if the old token is within
		// 15 mins of expiry.
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 1*time.Minute {
			// Gets the refresh token from the cookie.
			// Parses token and checks if it valid.
			tkn, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					c.Response().Writer.WriteHeader(http.StatusUnauthorized)
				}
				log.Debugf("Claim could not be parsed, %v", err)
			}

			if tkn != nil && tkn.Valid {
				// If everything is good, update tokens.
				token, err := CreateJWTToken(claims.UserId)
				if err != nil {
					c.Response().Writer.WriteHeader(http.StatusUnauthorized)
				} else {
					c.Response().Header().Set(headerName, token)
				}
			}
		}

		return next(c)
	}
}
