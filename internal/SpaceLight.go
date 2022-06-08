package main

import (
	"os"

	"github.com/BeanCodeDe/SpaceLight-AuthMiddleware/authAdapter"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/api"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/db"
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
	setLogLevel(os.Getenv("LOG_LEVEL"))
	log.Info("Start Server")
	db.Init()
	err := authAdapter.Init()
	if err != nil {
		log.Fatalf("Error while init authAdapter: %v", err)
	}
	e := echo.New()
	e.HTTPErrorHandler = api.CustomHTTPErrorHandler
	e.Validator = &CustomValidator{validator: validator.New()}
	profilGroup := e.Group(api.ProfilRootPath)
	profilGroup.Use(authAdapter.AuthMiddleware)
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
