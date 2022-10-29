package core

import (
	"fmt"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/db"
	"github.com/google/uuid"
)

type (
	HangarCore struct {
		Id     uuid.UUID
		UserId uuid.UUID
		Ship   *ShipCore
	}
)

func CreateDefaultHangar(userId uuid.UUID) error {
	hangar := &HangarCore{Id: uuid.New(), UserId: userId}
	err := hangar.mapToHangarDB().Create()
	if err != nil {
		return fmt.Errorf("error while creating hangar: %v", err)
	}
	return nil
}

func LoadAllHangars(userId uuid.UUID) ([]*HangarCore, error) {
	hangarDBList, err := db.GetHangarsByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("error while loading hangar: %v", err)
	}
	hangarList := make([]*HangarCore, len(hangarDBList))
	for index, hangarDB := range hangarDBList {
		shipDB, err := db.GetShipById(hangarDB.ShipId)
		if err != nil {
			return nil, fmt.Errorf("error while loading ship for hangar: %v", err)
		}
		hangarList[index] = mapToHangarCore(hangarDB, shipDB)
	}
	return hangarList, nil
}

func (hangar *HangarCore) mapToHangarDB() *db.HangarDB {
	return &db.HangarDB{Id: hangar.Id, UserId: hangar.UserId, ShipId: hangar.Ship.ShipId}
}

func mapToHangarCore(dbHangar *db.HangarDB, dbShip *db.ShipDB) *HangarCore {
	return &HangarCore{Id: dbHangar.Id, UserId: dbHangar.UserId, Ship: mapToShipCore(dbShip)}
}
