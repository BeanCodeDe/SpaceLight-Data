package api

import (
	"SpaceLight/internal/core"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const UserRootPath = "/user"

type (
	user interface {
		mapToUserCore() *core.UserCore
	}
	userCreateDTO struct {
		Name     string `json:"Name" validate:"required,alphanum"`
		Password string `json:"Password" validate:"required"`
	}

	userLoginDTO struct {
		ID       uuid.UUID `json:"ID"`
		Name     string    `json:"Name" validate:"required,alphanum"`
		Password string    `json:"Password" validate:"required"`
	}

	userResponseDTO struct {
		ID        uuid.UUID
		Name      string
		CreatedOn time.Time
		LastLogin time.Time
	}
)

func InitUserInterface(group *echo.Group) {
	group.GET("/login", login)
	group.PUT("", create)
}

func login(context echo.Context) error {
	log.Debugf("Login some user")
	userCore, err := bind(context, new(userLoginDTO))
	if err != nil {
		return err
	}

	token, err := userCore.Login()
	if err != nil {
		return err
	}

	log.Debugf("Logged in user %s, %s", userCore.ID, userCore.Name)
	return context.String(http.StatusOK, token)
}

func create(context echo.Context) error {
	log.Debugf("Create some user")
	userCore, err := bind(context, new(userCreateDTO))
	if err != nil {
		return err
	}
	createdUser, err := userCore.Create()
	if err != nil {
		return err
	}

	log.Debugf("Created user %s, %s", createdUser.ID, createdUser.Name)
	userResponseDTO := mapToUserResponseDTO(createdUser)
	return context.JSON(http.StatusCreated, userResponseDTO)
}

func (user *userCreateDTO) mapToUserCore() *core.UserCore {
	return &core.UserCore{Name: user.Name, Password: user.Password}
}

func (user *userLoginDTO) mapToUserCore() *core.UserCore {
	return &core.UserCore{ID: user.ID, Name: user.Name, Password: user.Password}
}

func mapToUserResponseDTO(user *core.UserCore) *userResponseDTO {
	return &userResponseDTO{ID: user.ID, Name: user.Name, CreatedOn: user.CreatedOn, LastLogin: user.LastLogin}
}

func bind(context echo.Context, toBindUser user) (*core.UserCore, error) {
	log.Debugf("Bind context to user %v", context)
	if err := context.Bind(toBindUser); err != nil {
		log.Warnf("Could not bind user, %v", err)
		return nil, echo.ErrBadRequest
	}
	log.Debugf("User bind %v", toBindUser)
	if err := context.Validate(toBindUser); err != nil {
		log.Warnf("Could not validate user, %v", err)
		return nil, echo.ErrBadRequest
	}
	return toBindUser.mapToUserCore(), nil
}
