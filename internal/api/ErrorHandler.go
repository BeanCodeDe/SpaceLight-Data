package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Warnf("An Error accurd: %v", err)
	c.String(http.StatusUnauthorized, "")
}
