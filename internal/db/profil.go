package db

import (
	"context"
	"errors"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/dataErr"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	log "github.com/sirupsen/logrus"
)

type ProfilDB struct {
	UserId uuid.UUID `db:"user_id"`
	Name   string    `db:"name"`
}

func (profil *ProfilDB) Create() error {
	log.Debugf("Create profil")

	if _, err := getConnection().Exec(context.Background(), "INSERT INTO spacelight.profil(user_id, name) VALUES($1, $2)", profil.UserId, profil.Name); err != nil {
		log.Errorf("Unknown error when inserting profil: %v", err)
		return dataErr.UnknownError
	}
	log.Debugf("Profil inserted into database")
	return nil
}

func GetProfilById(userId uuid.UUID) (*ProfilDB, error) {
	log.Debugf("Get profil by UserId %s", userId)

	var profils []*ProfilDB
	if err := pgxscan.Select(context.Background(), getConnection(), &profils, `SELECT user_id, name FROM spacelight.profil WHERE user_id = $1`, userId); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.NoDataFound:
				log.Warnf("No Profil for user with id %s found", userId)
				return nil, dataErr.ProfilForUserNotFoundError
			}
		}
		log.Errorf("Unknown error when getting profil by id: %v", err)
		return nil, dataErr.UnknownError
	}

	if len(profils) != 1 {
		log.Errorf("Cant find only one profil. ProfilLists: %v", profils)
		return nil, dataErr.UnknownError
	}

	log.Debugf("Got profil %v", profils[0])
	return profils[0], nil
}
