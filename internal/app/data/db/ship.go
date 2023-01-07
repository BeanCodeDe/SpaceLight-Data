package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

type (
	ShipType struct {
		Id              uuid.UUID      `db:"id"`
		Name            string         `db:"name"`
		RoomPlaceList   []*RoomPlace   `db:"rome_place_list"`
		DoorList        []*Door        `db:"door_list"`
		WeaponPlaceList []*WeaponPlace `db:"weapon_place_list"`
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
		Type           string
		RoomBlockOneId int
		RoomBlockTwoId int
	}

	WeaponPlace struct {
		PosX int
		PosY int
	}
)

func (db *postgresConnection) GetAllShipTypes() ([]*ShipType, error) {
	var shipTypes []*ShipType
	if err := pgxscan.Select(context.Background(), db.dbPool, &shipTypes, `SELECT id, name, rome_place_list, door_list, weapon_place_list FROM spacelight_data.ship_type`); err != nil {
		return nil, fmt.Errorf("error while selecting all ship types: %w", err)
	}

	return shipTypes, nil
}
