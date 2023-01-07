package db

import (
	"errors"
	"strings"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/util"
	"github.com/google/uuid"
)

type (
	DB interface {
		Close()
		CreateProfile(profile *ProfileDB) error
		GetProfileById(userId uuid.UUID) (*ProfileDB, error)
		GetAllShipTypes() ([]*ShipType, error)
	}
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func NewConnection() (DB, error) {
	switch db := strings.ToLower(util.GetEnvWithFallback("DATABASE", "postgresql")); db {
	case "postgresql":
		return newPostgresConnection()
	default:
		return nil, errors.New("no configuration for %s found")
	}
}
