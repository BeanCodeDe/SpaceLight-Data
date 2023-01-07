package api

import (
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/core"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const ShipRootPath = "/ships"

type (
	ShipTypeDTO struct {
		Id              uuid.UUID
		Name            string
		RoomPlaceList   []*RoomPlaceDTO
		DoorList        []*DoorDTO
		WeaponPlaceList []*WeaponPlaceDTO
	}

	RoomPlaceDTO struct {
		PosX int
		PosY int
	}

	DoorDTO struct {
		Type        string
		Orientation string
		PosX        int
		PosY        int
	}

	WeaponPlaceDTO struct {
		PosX int
		PosY int
	}
)

func initShipsInterface(group *echo.Group, api *EchoApi) {
	group.GET("", api.getAllShipTypes)
}

func (api *EchoApi) getAllShipTypes(context echo.Context) error {
	shipTypeList, err := api.core.GetAllShipTypes()
	if err != nil {
		log.Warnf("Error while loading ship types: %w", err)
		return echo.ErrInternalServerError
	}

	mappedShipTypeList := make([]*ShipTypeDTO, len(shipTypeList))
	for index, shipType := range shipTypeList {
		mappedShipTypeList[index] = mapToShipType(shipType)
	}

	return context.JSON(http.StatusOK, mappedShipTypeList)
}

func mapToShipType(shipType *core.ShipType) *ShipTypeDTO {
	roomPlaceList := make([]*RoomPlaceDTO, len(shipType.RoomPlaceList))
	for index, roomPlace := range shipType.RoomPlaceList {
		roomPlaceList[index] = mapToRoomPlace(roomPlace)
	}

	doorList := make([]*DoorDTO, len(shipType.DoorList))
	for index, door := range shipType.DoorList {
		doorList[index] = mapToDoor(door)
	}

	weaponPlaceList := make([]*WeaponPlaceDTO, len(shipType.WeaponPlaceList))
	for index, weaponPlace := range shipType.WeaponPlaceList {
		weaponPlaceList[index] = mapToWeaponPlace(weaponPlace)
	}

	return &ShipTypeDTO{
		Id:              shipType.Id,
		Name:            shipType.Name,
		RoomPlaceList:   roomPlaceList,
		DoorList:        doorList,
		WeaponPlaceList: weaponPlaceList,
	}
}

func mapToRoomPlace(roomPlace *core.RoomPlace) *RoomPlaceDTO {
	return &RoomPlaceDTO{PosX: roomPlace.PosX, PosY: roomPlace.PosY}
}

func mapToDoor(door *core.Door) *DoorDTO {
	return &DoorDTO{
		Type:        door.Type,
		Orientation: door.Orientation,
		PosX:        door.PosX,
		PosY:        door.PosY,
	}
}

func mapToWeaponPlace(weaponPlace *core.WeaponPlace) *WeaponPlaceDTO {
	return &WeaponPlaceDTO{PosX: weaponPlace.PosX, PosY: weaponPlace.PosY}
}
