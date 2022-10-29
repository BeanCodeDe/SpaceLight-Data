package core

import (
	"fmt"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/adapter"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/db"
	"github.com/google/uuid"
)

type (
	ProfilCore struct {
		UserId uuid.UUID
		Name   string
	}
)

func (profil *ProfilCore) Create(password string) (*ProfilCore, error) {
	userId, err := adapter.CreateUser(password)
	if err != nil {
		return nil, fmt.Errorf("something went wrong, while creating user in auth service: %v", err)
	}
	profil.UserId = userId
	err = profil.mapToProfilDB().Create()
	if err != nil {
		return nil, fmt.Errorf("something went wrong, while persisting profil: %v", err)
	}
	dbProfil, err := db.GetProfilById(profil.UserId)
	if err != nil {
		return nil, fmt.Errorf("something went wrong, while getting profil: %v", err)
	}
	return mapToUserCore(dbProfil), nil
}

func LoadProfil(userId uuid.UUID) (*ProfilCore, error) {
	dbProfil, err := db.GetProfilById(userId)
	if err != nil {
		return nil, fmt.Errorf("something went wrong, while getting profil: %v", err)
	}
	return mapToUserCore(dbProfil), nil
}

func (profil *ProfilCore) mapToProfilDB() *db.ProfilDB {
	return &db.ProfilDB{UserId: profil.UserId, Name: profil.Name}
}

func mapToUserCore(dbProfil *db.ProfilDB) *ProfilCore {
	return &ProfilCore{UserId: dbProfil.UserId, Name: dbProfil.Name}
}
