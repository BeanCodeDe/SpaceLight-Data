package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type HangarDB struct {
	Id     uuid.UUID `db:"id"`
	ShipId uuid.UUID `db:"ship_id"`
	UserId uuid.UUID `db:"user_id"`
}

func (hangar *HangarDB) Create() error {
	if _, err := getConnection().Exec(context.Background(), "INSERT INTO spacelight.hangar(id, ship_id, user_id) VALUES($1, $2, $3)", hangar.Id, hangar.ShipId, hangar.UserId); err != nil {
		return fmt.Errorf("unknown error when inserting profil: %v", err)
	}
	return nil
}

func GetHangarsByUserId(userId uuid.UUID) ([]*HangarDB, error) {
	var hangarList []*HangarDB
	if err := pgxscan.Select(context.Background(), getConnection(), &hangarList, `SELECT id, ship_id, user_id FROM spacelight.hangar WHERE user_id = $1`, userId); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.NoDataFound:
				return nil, fmt.Errorf("no hangar for user with id %s found", userId)
			}
		}
		return nil, fmt.Errorf("unknown error when getting hangars by user id: %v", err)
	}

	return hangarList, nil
}
