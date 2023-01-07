package api

import (
	"fmt"
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/core"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const ProfilRootPath = "/profiles"

type (
	profileCreateDTO struct {
		Name string `json:"Name" validate:"required,alphanum"`
	}

	profileResponseDTO struct {
		UserId uuid.UUID
		Name   string
	}
)

func initProfilInterface(group *echo.Group, api *EchoApi) {
	group.GET("/:"+userIdParam, api.getProfil)
	group.PUT("/:"+userIdParam, api.createProfil)
}

func (api *EchoApi) getProfil(context echo.Context) error {
	log.Debugf("Get Profile")
	userId, err := getUserId(context)
	if err != nil {
		log.Warnf("Error while getting user id: %w", err)
		return echo.ErrBadRequest
	}
	if err := checkUserId(userId, context); err != nil {
		log.Warnf("Error while checking user id: %w", err)
		return echo.ErrUnauthorized
	}

	profil, err := api.core.GetProfile(userId)
	if err != nil {
		log.Warnf("Error while loading profil: %w", err)
		return echo.ErrInternalServerError
	}

	profilResponseDTO := &profileResponseDTO{UserId: profil.UserId, Name: profil.Name}
	return context.JSON(http.StatusOK, profilResponseDTO)
}

func (api *EchoApi) createProfil(context echo.Context) error {
	log.Debugf("Create some profil")
	userId, err := getUserId(context)
	if err != nil {
		log.Warnf("Error while getting user id: %w", err)
		return echo.ErrBadRequest
	}

	if err := checkUserId(userId, context); err != nil {
		log.Warnf("Error while checking user id: %w", err)
		return echo.ErrUnauthorized
	}

	profileCore, err := bindCreationDTO(context, new(profileCreateDTO))
	if err != nil {
		log.Warnf("Error while mapping profil: %v", err)
		return echo.ErrBadRequest
	}

	if err := api.core.CreateProfile(profileCore); err != nil {
		log.Warnf("Error while creating profil: %v", err)
		return echo.ErrInternalServerError
	}

	log.Debugf("Created profile %s with name %s", profileCore.UserId, profileCore.Name)
	return context.NoContent(http.StatusCreated)
}

func bindCreationDTO(context echo.Context, toBindProfil *profileCreateDTO) (*core.Profile, error) {
	if err := context.Bind(toBindProfil); err != nil {
		return nil, fmt.Errorf("could not bind profil, %w", err)
	}
	if err := context.Validate(toBindProfil); err != nil {
		return nil, fmt.Errorf("could not validate profil, %w", err)
	}
	return &core.Profile{Name: toBindProfil.Name}, nil
}
