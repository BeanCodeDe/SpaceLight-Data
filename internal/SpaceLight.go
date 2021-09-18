package main

import (
	"SpaceLight/internal/api"
	"SpaceLight/internal/db"

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
	log.Info("Start Server")
	db.Init()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	//e.Use(GetJWTConfig())
	userGroup := e.Group(api.UserRootPath)
	api.InitUserInterface(userGroup)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleExit() {
	db.Close()
}
