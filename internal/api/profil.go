package api

import (
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-AuthMiddleware/authAdapter"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/core"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const ProfilRootPath = "/profil"

type (
	profilCreateDTO struct {
		Name     string `json:"Name" validate:"required,alphanum"`
		Password string `json:"Password" validate:"required"`
	}

	profilResponseDTO struct {
		UserId uuid.UUID
		Name   string
	}
)

func InitProfilInterface(group *echo.Group) {
	group.GET("", getProfil, authAdapter.CheckToken)
	group.PUT("", createProfil)
}

func getProfil(context echo.Context) error {
	log.Debugf("Get Profile")
	return context.String(http.StatusOK, "")
}

func createProfil(context echo.Context) error {
	log.Debugf("Create some profil")
	profilCore, password, err := bindCreationDTO(context, new(profilCreateDTO))
	if err != nil {
		return err
	}
	createdProfil, err := profilCore.Create(password)
	if err != nil {
		return err
	}

	log.Debugf("Created profile %s with name %s", createdProfil.UserId, createdProfil.Name)
	profilResponseDTO := mapToProfilResponseDTO(createdProfil)
	return context.JSON(http.StatusCreated, profilResponseDTO)
}

func (profil *profilCreateDTO) mapToUserCore() *core.ProfilCore {
	return &core.ProfilCore{Name: profil.Name}
}

func mapToProfilResponseDTO(profil *core.ProfilCore) *profilResponseDTO {
	return &profilResponseDTO{UserId: profil.UserId, Name: profil.Name}
}

func bindCreationDTO(context echo.Context, toBindProfil *profilCreateDTO) (*core.ProfilCore, string, error) {
	log.Debugf("Bind context to profil %v", context)
	if err := context.Bind(toBindProfil); err != nil {
		log.Warnf("Could not bind profil, %v", err)
		return nil, "", echo.ErrBadRequest
	}
	log.Debugf("User bind %v", toBindProfil)
	if err := context.Validate(toBindProfil); err != nil {
		log.Warnf("Could not validate profil, %v", err)
		return nil, "", echo.ErrBadRequest
	}
	return toBindProfil.mapToUserCore(), toBindProfil.Password, nil
}
