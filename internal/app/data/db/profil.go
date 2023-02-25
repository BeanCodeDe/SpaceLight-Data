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

var (
	ErrProfileAlreadyExists = errors.New("profile already exists")
)

type (
	ProfileDB struct {
		UserId uuid.UUID `db:"user_id"`
		Name   string    `db:"name"`
	}
)

func (db *postgresConnection) CreateProfile(profile *ProfileDB) error {
	if _, err := db.dbPool.Exec(context.Background(), "INSERT INTO spacelight_data.profile(user_id, name) VALUES($1, $2)", profile.UserId, profile.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return ErrProfileAlreadyExists
			}
		}

		return fmt.Errorf("unknown error when inserting user: %v", err)
	}
	return nil
}

func (db *postgresConnection) GetProfileById(userId uuid.UUID) (*ProfileDB, error) {
	var profiles []*ProfileDB
	if err := pgxscan.Select(context.Background(), db.dbPool, &profiles, `SELECT user_id, name FROM spacelight_data.profile WHERE user_id = $1`, userId); err != nil {
		return nil, fmt.Errorf("error while selecting user with id %v: %w", userId, err)
	}

	if len(profiles) != 1 {
		return nil, fmt.Errorf("cant find only one profile. ProfilLists: %v", profiles)
	}

	return profiles[0], nil
}
