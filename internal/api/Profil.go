package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const ProfilRootPath = "/profil"

func InitProfilInterface(group *echo.Group) {
	group.GET("", profil)
}

func profil(context echo.Context) error {
	log.Debugf("Get Profile")

	return context.String(http.StatusOK, "")
}
