package core

import (
	"SpaceLight/internal/db"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	UserCore struct {
		ID        uuid.UUID
		Name      string
		Password  string
		CreatedOn time.Time
		LastLogin time.Time
	}
)

func (user *UserCore) Create() (createdUser *UserCore, err error) {
	log.Debugf("Create user %s, %s", user.Name)
	if err = user.mapToUserDB().Create(); err != nil {
		return nil, err
	}

	dbUser, err := db.GetUserByName(user.Name)
	if err != nil {
		return nil, err
	}

	return mapToUserCore(dbUser), nil
}

func (user *UserCore) Login() (string, error) {
	log.Debugf("Login user %s, %s", user.ID, user.Name)
	logedIn, err := db.CheckPassword(user.ID, user.Name, user.Password)
	if err != nil {
		log.Warnf("Could not check password, %v", err)
		return "", err
	}
	if !logedIn {
		log.Debugf("Wrong auth data, %v", user)
		return "", echo.ErrUnauthorized
	}

	return createJWTToken(user.ID)
}

func (user *UserCore) mapToUserDB() *db.UserDB {
	return &db.UserDB{ID: user.ID, Name: user.Name, Password: user.Password, CreatedOn: user.CreatedOn, LastLogin: user.LastLogin}
}

func mapToUserCore(dbUser *db.UserDB) *UserCore {
	return &UserCore{ID: dbUser.ID, Name: dbUser.Name, Password: dbUser.Password, CreatedOn: dbUser.CreatedOn, LastLogin: dbUser.LastLogin}
}
