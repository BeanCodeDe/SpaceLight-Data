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
		Id            int
		RoomBlockList []*RoomBlockDTO
	}

	RoomBlockDTO struct {
		Id   int
		PosX int
		PosY int
	}

	DoorDTO struct {
		RoomBlockOneId int
		RoomBlockTwoId int
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
		log.Warnf("Error while loading ship types: %v", err)
		return echo.ErrInternalServerError
	}

	mappedShipTypeList := make([]*ShipTypeDTO, len(shipTypeList))
	for index, shipType := range shipTypeList {
		mappedShipTypeList[index] = mapToShipType(shipType)
	}

	return context.JSON(http.StatusOK, mappedShipTypeList)
}

func mapToShipType(shipType *core.ShipType) *ShipTypeDTO {

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
		RoomPlaceList:   mapToRoomPlaceList(shipType.RoomPlaceList),
		DoorList:        doorList,
		WeaponPlaceList: weaponPlaceList,
	}
}

func mapToRoomPlaceList(coreRoomPlaceList []*core.RoomPlace) []*RoomPlaceDTO {
	roomPlaceList := make([]*RoomPlaceDTO, len(coreRoomPlaceList))
	for index, roomPlace := range coreRoomPlaceList {
		roomPlaceList[index] = mapToRoomPlace(roomPlace)
	}
	return roomPlaceList
}

func mapToRoomPlace(roomPlace *core.RoomPlace) *RoomPlaceDTO {
	return &RoomPlaceDTO{Id: roomPlace.Id, RoomBlockList: mapToRoomBlockList(roomPlace.RoomBlockList)}
}

func mapToRoomBlockList(coreRoomBlockList []*core.RoomBlock) []*RoomBlockDTO {
	roomBlockList := make([]*RoomBlockDTO, len(coreRoomBlockList))
	for index, roomBlock := range coreRoomBlockList {
		roomBlockList[index] = mapToRoomBlock(roomBlock)
	}
	return roomBlockList
}

func mapToRoomBlock(coreRoomBlock *core.RoomBlock) *RoomBlockDTO {
	return &RoomBlockDTO{Id: coreRoomBlock.Id, PosX: coreRoomBlock.PosX, PosY: coreRoomBlock.PosY}
}

func mapToDoorList(coreDoorList []*core.Door) []*DoorDTO {
	doorList := make([]*DoorDTO, len(coreDoorList))
	for index, door := range coreDoorList {
		doorList[index] = mapToDoor(door)
	}
	return doorList
}

func mapToDoor(door *core.Door) *DoorDTO {
	return &DoorDTO{
		RoomBlockOneId: door.RoomBlockOneId,
		RoomBlockTwoId: door.RoomBlockTwoId,
	}
}

func mapToWeaponPlace(weaponPlace *core.WeaponPlace) *WeaponPlaceDTO {
	return &WeaponPlaceDTO{PosX: weaponPlace.PosX, PosY: weaponPlace.PosY}
}
