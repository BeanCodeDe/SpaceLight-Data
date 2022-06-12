package core

import (
	"github.com/BeanCodeDe/SpaceLight-Data/internal/adapter"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/db"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type (
	ProfilCore struct {
		UserId uuid.UUID
		Name   string
	}
)

func (profil *ProfilCore) Create(password string) (*ProfilCore, error) {
	log.Debug("Create Profil")
	userId, err := adapter.CreateUser(password)
	if err != nil {
		log.Errorf("Something went wrong, while creating user in auth service: %v", err)
		return nil, err
	}
	profil.UserId = userId
	err = profil.mapToProfilDB().Create()
	if err != nil {
		log.Errorf("Something went wrong, while persisting profil: %v", err)
		return nil, err
	}
	dbProfil, err := db.GetProfilById(profil.UserId)
	if err != nil {
		log.Errorf("Something went wrong, while getting profil: %v", err)
		return nil, err
	}
	return mapToUserCore(dbProfil), nil
}

func (profil *ProfilCore) mapToProfilDB() *db.ProfilDB {
	return &db.ProfilDB{UserId: profil.UserId, Name: profil.Name}
}

func mapToUserCore(dbProfil *db.ProfilDB) *ProfilCore {
	return &ProfilCore{UserId: dbProfil.UserId, Name: dbProfil.Name}
}
