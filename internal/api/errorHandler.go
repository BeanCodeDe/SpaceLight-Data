package api

import (
	"github.com/BeanCodeDe/SpaceLight-Data/internal/dataErr"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Warnf("An Error accurd: %v", err)

	echoError, ok := err.(*echo.HTTPError)
	if ok {
		c.String(echoError.Code, "")
		return
	}

	customError, ok := err.(*dataErr.CustomError)
	if ok {
		c.String(customError.HttpCode, "")
		return
	}

	c.String(echo.ErrUnauthorized.Code, "")
}