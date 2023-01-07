package core

import (
	"fmt"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/db"
)

func (core *CoreFacade) GetAllShipTypes() ([]*ShipType, error) {
	shipTypeList, err := core.db.GetAllShipTypes()
	if err != nil {
		return nil, fmt.Errorf("error while getting all ship types from database: %w", err)
	}

	mappedShipTypeList := make([]*ShipType, len(shipTypeList))
	for index, shipType := range shipTypeList {
		mappedShipTypeList[index] = mapToShipType(shipType)
	}

	return mappedShipTypeList, nil
}

func mapToShipType(shipType *db.ShipType) *ShipType {
	roomPlaceList := make([]*RoomPlace, len(shipType.RoomPlaceList))
	for index, roomPlace := range shipType.RoomPlaceList {
		roomPlaceList[index] = mapToRoomPlace(roomPlace)
	}

	doorList := make([]*Door, len(shipType.DoorList))
	for index, door := range shipType.DoorList {
		doorList[index] = mapToDoor(door)
	}

	weaponPlaceList := make([]*WeaponPlace, len(shipType.WeaponPlaceList))
	for index, weaponPlace := range shipType.WeaponPlaceList {
		weaponPlaceList[index] = mapToWeaponPlace(weaponPlace)
	}

	return &ShipType{
		Id:              shipType.Id,
		Name:            shipType.Name,
		RoomPlaceList:   roomPlaceList,
		DoorList:        doorList,
		WeaponPlaceList: weaponPlaceList,
	}
}

func mapToRoomPlace(roomPlace *db.RoomPlace) *RoomPlace {
	return &RoomPlace{PosX: roomPlace.PosX, PosY: roomPlace.PosY}
}

func mapToDoor(door *db.Door) *Door {
	return &Door{
		Type:        door.Type,
		Orientation: door.Orientation,
		PosX:        door.PosX,
		PosY:        door.PosY,
	}
}

func mapToWeaponPlace(weaponPlace *db.WeaponPlace) *WeaponPlace {
	return &WeaponPlace{PosX: weaponPlace.PosX, PosY: weaponPlace.PosY}
}
