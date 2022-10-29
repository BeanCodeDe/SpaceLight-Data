package db

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type ShipDB struct {
	ShipId   uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	RoomList []*RoomDB `db:"room_list"`
	DoorList []*DoorDB `db:"door_list"`
}

type RoomDB struct {
	RoomTyp string
	PosX    int
	PosY    int
}

type DoorDB struct {
	RoomOnePosX int
	RoomOnePosY int
	RoomTwoPosX int
	RoomTwoPosY int
}

func GetShipById(shipId uuid.UUID) (*ShipDB, error) {
	var shipArray []*ShipDB
	if err := pgxscan.Select(context.Background(), getConnection(), &shipArray, `SELECT id, name, room_list, door_list FROM spacelight.ship WHERE id = $1`, shipId); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.NoDataFound:
				return nil, fmt.Errorf("no ship found for id %s", shipId)
			}
		}
		return nil, fmt.Errorf("unknown error when getting profil by id: %v", err)
	}

	if len(shipArray) != 1 {
		return nil, fmt.Errorf("cant find only one ship. Shiplist: %v", shipArray)
	}

	return shipArray[0], nil
}

func (room RoomDB) Value() (driver.Value, error) {
	return json.Marshal(room)
}

func (room *RoomDB) Scan(value interface{}) error {
	if b, ok := value.([]byte); ok {
		return json.Unmarshal(b, room)
	}
	return nil
}

func (door DoorDB) Value() (driver.Value, error) {
	return json.Marshal(door)
}

func (door *DoorDB) Scan(value interface{}) error {
	if b, ok := value.([]byte); ok {
		return json.Unmarshal(b, door)
	}
	return nil
}
