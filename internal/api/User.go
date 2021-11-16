package api

import (
	"SpaceLight/internal/core"

	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const UserRootPath = "/user"

func InitUserInterface(group *echo.Group) {
	group.GET("", login)
	group.PUT("", Create)

}

func login(context echo.Context) error {
	user := new(core.UserDTO)
	logedIn, err := user.CheckPassword()
	if err != nil {
		return err
	}
	if !logedIn {
		return echo.ErrUnauthorized
	}
	token, err := createJWTToken(user.Name)
	if err != nil {
		return err
	}
	return context.JSON(200, token)
}

func Create(context echo.Context) error {
	user := new(core.UserDTO)
	log.Debugf("create user %s", user.Name)
	if err := context.Bind(user); err != nil {
		log.Warnf("Could not bind user ,%v", err)
		return echo.ErrBadRequest
	}
	if err := context.Validate(user); err != nil {
		log.Warnf("Could not validate user ,%v", err)
		return echo.ErrBadRequest
	}
	createdUser, err := user.Create()
	if err != nil {
		return err
	}

	log.Debugf("created user %v", createdUser)
	return context.JSON(201, createdUser)
}
