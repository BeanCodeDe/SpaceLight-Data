package api

import (
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-AuthMiddleware/authAdapter"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/core"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

const HangarRootPath = "/hangar"

type (
	hangarDTO struct {
		Id     uuid.UUID
		UserId uuid.UUID
		Ship   *shipDTO
	}
)

func InitHangarInterface(group *echo.Group) {
	group.GET("", getHangars, authAdapter.CheckToken)
}

func getHangars(context echo.Context) error {
	log.Debugf("Get Hangar")
	claims := context.Get(authAdapter.ClaimName).(*authAdapter.Claims)

	hangarCoreList, err := core.LoadAllHangars(claims.UserId)
	if err != nil {
		log.Warnf("Error while loading hangars: %v", err)
		return echo.ErrInternalServerError
	}

	hangarDTOList := make([]*hangarDTO, len(hangarCoreList))
	for index, hangar := range hangarCoreList {
		hangarDTOList[index] = mapToHangarDTO(hangar)
	}
	return context.JSON(http.StatusOK, hangarDTOList)
}

func (hangar *hangarDTO) mapToHangarCore() *core.HangarCore {
	return &core.HangarCore{Id: hangar.Id, UserId: hangar.UserId, Ship: hangar.Ship.mapToShipCore()}
}

func mapToHangarDTO(hangar *core.HangarCore) *hangarDTO {
	return &hangarDTO{Id: hangar.Id, UserId: hangar.UserId, Ship: mapToShipDTO(hangar.Ship)}
}
