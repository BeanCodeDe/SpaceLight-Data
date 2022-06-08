package api

import (
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-AuthMiddleware/authAdapter"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const ProfilRootPath = "/profil"

func InitProfilInterface(group *echo.Group) {
	group.GET("", profil)
}

func profil(context echo.Context) error {
	log.Debugf("Get Profile: %v", context)
	log.Debugf("header %s", context.Response().Header().Get(authAdapter.AuthName))
	return context.String(http.StatusOK, "")
}
