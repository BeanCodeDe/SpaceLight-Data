package main

import (
	"github.com/BeanCodeDe/SpaceLight-AuthMiddleware/authAdapter"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/api"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/config"
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
	setLogLevel(config.LogLevel)
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
	api.InitProfilInterface(profilGroup)
	hangarGroup := e.Group(api.HangarRootPath)
	api.InitHangarInterface(hangarGroup)
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
	default:
		log.SetLevel(log.DebugLevel)
		log.Errorf("Log level %s unknow", logLevel)
	}

}

func handleExit() {
	db.Close()
}
