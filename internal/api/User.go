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
	return context.JSON(200, "jwtToken")
}

func Create(context echo.Context) error {
	user := new(core.UserDTO)
	log.Infof("create user %s", user.UserName)
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

	log.Infof("created user %v", createdUser)
	return context.JSON(201, createdUser)
}
