package core

import (
	"fmt"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/db"
)

func (core *CoreFacade) GetAllShipTypes() ([]*ShipType, error) {
	shipTypeList, err := core.db.GetAllShipTypes()
	if err != nil {
		return nil, fmt.Errorf("error while getting all ship types from database: %v", err)
	}

	mappedShipTypeList := make([]*ShipType, len(shipTypeList))
	for index, shipType := range shipTypeList {
		mappedShipTypeList[index] = mapToShipType(shipType)
	}

	return mappedShipTypeList, nil
}

func mapToShipType(shipType *db.ShipType) *ShipType {

	return &ShipType{
		Id:              shipType.Id,
		Name:            shipType.Name,
		RoomPlaceList:   mapToRoomPlaceList(shipType.RoomPlaceList),
		DoorList:        mapToDoorList(shipType.DoorList),
		WeaponPlaceList: mapToWeaponPlaceList(shipType.WeaponPlaceList),
	}
}

func mapToRoomPlaceList(dbRoomPlaceList []*db.RoomPlace) []*RoomPlace {
	roomPlaceList := make([]*RoomPlace, len(dbRoomPlaceList))
	for index, roomPlace := range dbRoomPlaceList {
		roomPlaceList[index] = mapToRoomPlace(roomPlace)
	}
	return roomPlaceList
}

func mapToRoomPlace(roomPlace *db.RoomPlace) *RoomPlace {
	return &RoomPlace{Id: roomPlace.Id, RoomBlockList: mapToRoomBlockList(roomPlace.RoomBlockList)}
}

func mapToRoomBlockList(dbRoomBlockList []*db.RoomBlock) []*RoomBlock {
	roomBlockList := make([]*RoomBlock, len(dbRoomBlockList))
	for index, roomBlock := range dbRoomBlockList {
		roomBlockList[index] = mapToRoomBlock(roomBlock)
	}
	return roomBlockList
}

func mapToRoomBlock(dbRoomBlock *db.RoomBlock) *RoomBlock {
	return &RoomBlock{Id: dbRoomBlock.Id, PosX: dbRoomBlock.PosX, PosY: dbRoomBlock.PosY}
}

func mapToDoorList(dbDoorList []*db.Door) []*Door {
	doorList := make([]*Door, len(dbDoorList))
	for index, door := range dbDoorList {
		doorList[index] = mapToDoor(door)
	}
	return doorList
}

func mapToDoor(door *db.Door) *Door {
	return &Door{
		RoomBlockOneId: door.RoomBlockOneId,
		RoomBlockTwoId: door.RoomBlockTwoId,
	}
}

func mapToWeaponPlaceList(dbWeaponPlaceList []*db.WeaponPlace) []*WeaponPlace {
	weaponPlaceList := make([]*WeaponPlace, len(dbWeaponPlaceList))
	for index, weaponPlace := range dbWeaponPlaceList {
		weaponPlaceList[index] = mapToWeaponPlace(weaponPlace)
	}
	return weaponPlaceList
}

func mapToWeaponPlace(weaponPlace *db.WeaponPlace) *WeaponPlace {
	return &WeaponPlace{PosX: weaponPlace.PosX, PosY: weaponPlace.PosY}
}
