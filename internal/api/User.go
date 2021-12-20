package api

import (
	"SpaceLight/internal/core"
	"net/http"

	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const UserRootPath = "/user"

func InitUserInterface(group *echo.Group) {
	group.GET("/login", login)
	group.PUT("", Create)

}

func login(context echo.Context) error {
	user := new(core.UserLoginDTO)
	if err := context.Bind(user); err != nil {
		log.Warnf("Could not bind user, %v", err)
		return echo.ErrBadRequest
	}
	log.Debugf("login user %v", user)
	err := user.Login()
	if err != nil {
		return err
	}

	token, err := createJWTToken(user.Name)
	if err != nil {
		return err
	}
	log.Debugf("loged in user %v", user)
	return context.String(http.StatusOK, token)
}

func Create(context echo.Context) error {
	user := new(core.UserCreateDTO)
	if err := context.Bind(user); err != nil {
		log.Warnf("Could not bind user, %v", err)
		return echo.ErrBadRequest
	}
	log.Debugf("create user %v", user)
	if err := context.Validate(user); err != nil {
		log.Warnf("Could not validate user, %v", err)
		return echo.ErrBadRequest
	}
	createdUser, err := user.Create()
	if err != nil {
		return err
	}

	log.Debugf("created user %v", createdUser)
	return context.JSON(http.StatusCreated, createdUser)
}
