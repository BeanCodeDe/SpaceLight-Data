package db

import (
	"github.com/google/uuid"
)

type HangarDB struct {
	HangarId uuid.UUID `db:"id"`
	ShipId   uuid.UUID `db:"ship_id"`
	UserId   uuid.UUID `db:"user_id"`
}
