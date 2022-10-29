package core

import (
	"github.com/BeanCodeDe/SpaceLight-Data/internal/db"
	"github.com/google/uuid"
)

type (
	ShipCore struct {
		ShipId   uuid.UUID
		Name     string
		RoomList []*RoomCore
		DoorList []*DoorCore
	}

	RoomCore struct {
		RoomTyp string
		PosX    int
		PosY    int
	}

	DoorCore struct {
		RoomOnePosX int
		RoomOnePosY int
		RoomTwoPosX int
		RoomTwoPosY int
	}
)

func (ship *ShipCore) mapToShipDB() *db.ShipDB {
	roomDBList := make([]*db.RoomDB, len(ship.RoomList))
	for index, room := range ship.RoomList {
		roomDBList[index] = room.mapToRoomDB()
	}

	doorDBList := make([]*db.DoorDB, len(ship.DoorList))
	for index, door := range ship.DoorList {
		doorDBList[index] = door.mapToDoorDB()
	}
	return &db.ShipDB{ShipId: ship.ShipId, Name: ship.Name, RoomList: roomDBList, DoorList: doorDBList}
}

func (room *RoomCore) mapToRoomDB() *db.RoomDB {
	return &db.RoomDB{RoomTyp: room.RoomTyp, PosX: room.PosX, PosY: room.PosY}
}

func (door *DoorCore) mapToDoorDB() *db.DoorDB {
	return &db.DoorDB{RoomOnePosX: door.RoomOnePosX, RoomOnePosY: door.RoomOnePosY, RoomTwoPosX: door.RoomTwoPosX, RoomTwoPosY: door.RoomTwoPosY}
}

func mapToShipCore(ship *db.ShipDB) *ShipCore {
	roomCoreList := make([]*RoomCore, len(ship.RoomList))
	for index, room := range ship.RoomList {
		roomCoreList[index] = mapToRoomCore(room)
	}

	doorCoreList := make([]*DoorCore, len(ship.DoorList))
	for index, door := range ship.DoorList {
		doorCoreList[index] = mapToDoorCore(door)
	}
	return &ShipCore{ShipId: ship.ShipId, Name: ship.Name, RoomList: roomCoreList, DoorList: doorCoreList}
}

func mapToRoomCore(room *db.RoomDB) *RoomCore {
	return &RoomCore{RoomTyp: room.RoomTyp, PosX: room.PosX, PosY: room.PosY}
}

func mapToDoorCore(door *db.DoorDB) *DoorCore {
	return &DoorCore{RoomOnePosX: door.RoomOnePosX, RoomOnePosY: door.RoomOnePosY, RoomTwoPosX: door.RoomTwoPosX, RoomTwoPosY: door.RoomTwoPosY}
}
