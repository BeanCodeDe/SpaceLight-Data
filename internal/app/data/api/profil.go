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
	ProfileCreateDTO struct {
		Name string `json:"Name" validate:"required,alphanum"`
	}

	ProfileResponseDTO struct {
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
		log.Warnf("Error while getting user id: %v", err)
		return echo.ErrBadRequest
	}
	if err := checkUserId(userId, context); err != nil {
		log.Warnf("Error while checking user id: %v", err)
		return echo.ErrUnauthorized
	}

	profil, err := api.core.GetProfile(userId)
	if err != nil {
		log.Warnf("Error while loading profil: %v", err)
		return echo.ErrInternalServerError
	}

	profilResponseDTO := &ProfileResponseDTO{UserId: profil.UserId, Name: profil.Name}
	return context.JSON(http.StatusOK, profilResponseDTO)
}

func (api *EchoApi) createProfil(context echo.Context) error {
	log.Debugf("Create some profil")
	userId, err := getUserId(context)
	if err != nil {
		log.Warnf("Error while getting user id: %v", err)
		return echo.ErrBadRequest
	}

	if err := checkUserId(userId, context); err != nil {
		log.Warnf("Error while checking user id: %v", err)
		return echo.ErrUnauthorized
	}

	profileCore, err := bindCreationDTO(context, new(ProfileCreateDTO), userId)
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

func bindCreationDTO(context echo.Context, toBindProfil *ProfileCreateDTO, userId uuid.UUID) (*core.Profile, error) {
	if err := context.Bind(toBindProfil); err != nil {
		return nil, fmt.Errorf("could not bind profil, %v", err)
	}
	if err := context.Validate(toBindProfil); err != nil {
		return nil, fmt.Errorf("could not validate profil, %v", err)
	}
	return &core.Profile{Name: toBindProfil.Name, UserId: userId}, nil
}
