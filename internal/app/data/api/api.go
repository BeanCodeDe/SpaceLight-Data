package api

import (
	"fmt"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/core"
	"github.com/BeanCodeDe/authi/pkg/adapter"
	"github.com/BeanCodeDe/authi/pkg/middleware"
	"github.com/BeanCodeDe/authi/pkg/parser"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	correlationIdHeader = "X-Correlation-ID"
	userIdParam         = "userId"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
	EchoApi struct {
		core core.Core
	}
	Api interface {
	}
)

func NewApi() (Api, error) {
	core, err := core.NewCore()
	if err != nil {
		return nil, fmt.Errorf("error while creating core layer: %w", err)
	}

	authMiddleware, err := initAuthMiddleware()
	if err != nil {
		return nil, fmt.Errorf("error while creating auth middleware: %w", err)
	}

	echoApi := &EchoApi{core: core}
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(authMiddleware.CheckToken)

	profilesGroup := e.Group(ProfilRootPath)
	initProfilInterface(profilesGroup, echoApi)

	shipsGroup := e.Group(ShipRootPath)
	initShipsInterface(shipsGroup, echoApi)

	e.Logger.Fatal(e.Start(":1203"))
	return echoApi, nil
}

func initAuthMiddleware() (middleware.Middleware, error) {
	tokenParser, err := parser.NewJWTParser()
	if err != nil {
		return nil, fmt.Errorf("error while init auth middleware: %w", err)
	}
	return middleware.NewEchoMiddleware(tokenParser), nil
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func checkUserId(userId uuid.UUID, context echo.Context) error {
	claims := context.Get(adapter.ClaimName).(adapter.Claims)
	if claims.UserId != userId {
		return fmt.Errorf("user id of claim [%v] doesn't match user id from request [%v]", claims.UserId, userId)
	}
	return nil
}

func getUserId(context echo.Context) (uuid.UUID, error) {
	userId, err := uuid.Parse(context.Param(userIdParam))
	if err != nil {
		return uuid.Nil, fmt.Errorf("error while binding userId: %w", err)
	}
	return userId, nil
}
