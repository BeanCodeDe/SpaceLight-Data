package main

import (
	"SpaceLight/internal/api"
	"SpaceLight/internal/auth"
	"SpaceLight/internal/db"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo/v4"
)

//const rootPath = "/spacelight"

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	defer handleExit()
	setLogLevel(os.Getenv("SPACELIGHT_LOG_LEVEL"))
	log.Info("Start Server")
	db.Init()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	userGroup := e.Group(api.UserRootPath)
	api.InitUserInterface(userGroup)
	profilGroup := e.Group(api.ProfilRootPath)
	profilGroup.Use(auth.AuthMiddleware)
	api.InitProfilInterface(profilGroup)
	e.Logger.Fatal(e.Start(":1323"))
}

func setLogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	}
}

func handleExit() {
	db.Close()
}
