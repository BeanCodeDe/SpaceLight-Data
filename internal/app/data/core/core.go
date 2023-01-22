package core

import (
	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/db"
	"github.com/google/uuid"
)

type (
	//Profile
	Profile struct {
		UserId uuid.UUID
		Name   string
	}

	//Ship
	ShipType struct {
		Id              uuid.UUID
		Name            string
		RoomPlaceList   []*RoomPlace
		DoorList        []*Door
		WeaponPlaceList []*WeaponPlace
	}

	RoomPlace struct {
		Id            int
		RoomBlockList []*RoomBlock
	}

	RoomBlock struct {
		Id   int
		PosX int
		PosY int
	}

	Door struct {
		RoomBlockOneId int
		RoomBlockTwoId int
	}

	WeaponPlace struct {
		PosX int
		PosY int
	}

	//Facade
	CoreFacade struct {
		db db.DB
	}

	Core interface {
		CreateProfile(profileCore *Profile) error
		GetProfile(userId uuid.UUID) (*Profile, error)
		GetAllShipTypes() ([]*ShipType, error)
	}
)

func NewCore() (Core, error) {
	db, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	return &CoreFacade{db: db}, nil
}
