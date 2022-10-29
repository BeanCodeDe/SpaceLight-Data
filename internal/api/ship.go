package api

import (
	"github.com/BeanCodeDe/SpaceLight-Data/internal/core"
	"github.com/google/uuid"
)

const ShipRootPath = "/ship"

type (
	shipDTO struct {
		ShipId   uuid.UUID
		Name     string
		RoomList []*roomDTO
		DoorList []*doorDTO
	}

	roomDTO struct {
		RoomTyp string
		PosX    int
		PosY    int
	}

	doorDTO struct {
		RoomOnePosX int
		RoomOnePosY int
		RoomTwoPosX int
		RoomTwoPosY int
	}
)

func (ship *shipDTO) mapToShipCore() *core.ShipCore {
	roomCoreList := make([]*core.RoomCore, len(ship.RoomList))
	for index, room := range ship.RoomList {
		roomCoreList[index] = room.mapToRoomCore()
	}

	doorCoreList := make([]*core.DoorCore, len(ship.DoorList))
	for index, door := range ship.DoorList {
		doorCoreList[index] = door.mapToDoorCore()
	}
	return &core.ShipCore{ShipId: ship.ShipId, Name: ship.Name, RoomList: roomCoreList, DoorList: doorCoreList}
}

func (room *roomDTO) mapToRoomCore() *core.RoomCore {
	return &core.RoomCore{RoomTyp: room.RoomTyp, PosX: room.PosX, PosY: room.PosY}
}

func (door *doorDTO) mapToDoorCore() *core.DoorCore {
	return &core.DoorCore{RoomOnePosX: door.RoomOnePosX, RoomOnePosY: door.RoomOnePosY, RoomTwoPosX: door.RoomTwoPosX, RoomTwoPosY: door.RoomTwoPosY}
}

func mapToShipDTO(ship *core.ShipCore) *shipDTO {
	roomDTOList := make([]*roomDTO, len(ship.RoomList))
	for index, room := range ship.RoomList {
		roomDTOList[index] = mapToRoomDTO(room)
	}

	doorDTOList := make([]*doorDTO, len(ship.DoorList))
	for index, door := range ship.DoorList {
		doorDTOList[index] = mapToDoorDTO(door)
	}
	return &shipDTO{ShipId: ship.ShipId, Name: ship.Name, RoomList: roomDTOList, DoorList: doorDTOList}
}

func mapToRoomDTO(room *core.RoomCore) *roomDTO {
	return &roomDTO{RoomTyp: room.RoomTyp, PosX: room.PosX, PosY: room.PosY}
}

func mapToDoorDTO(door *core.DoorCore) *doorDTO {
	return &doorDTO{RoomOnePosX: door.RoomOnePosX, RoomOnePosY: door.RoomOnePosY, RoomTwoPosX: door.RoomTwoPosX, RoomTwoPosY: door.RoomTwoPosY}
}
