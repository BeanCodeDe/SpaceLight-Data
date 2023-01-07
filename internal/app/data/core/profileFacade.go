package core

import (
	"errors"
	"fmt"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/db"
	"github.com/google/uuid"
)

func (core *CoreFacade) CreateProfile(profile *Profile) error {
	profileDB := &db.ProfileDB{UserId: profile.UserId, Name: profile.Name}

	if err := core.db.CreateProfile(profileDB); err != nil {
		if errors.Is(err, db.ErrProfileAlreadyExists) {
			foundProfile, err := core.db.GetProfileById(profile.UserId)
			if err != nil {
				return fmt.Errorf("something went wrong while checking if user [%v] is already created: %w", profile.UserId, err)
			}

			if profile.Name == foundProfile.Name {
				return fmt.Errorf("profile name of request [%v] doesn't match name from database [%v]", profile.Name, foundProfile.Name)
			}

			return nil
		}
		return fmt.Errorf("error while creating profile: %v", err)
	}

	return nil
}

func (core *CoreFacade) GetProfile(userId uuid.UUID) (*Profile, error) {
	dbProfil, err := core.db.GetProfileById(userId)
	if err != nil {
		return nil, fmt.Errorf("something went wrong, while getting profil: %ws", err)
	}
	return &Profile{UserId: dbProfil.UserId, Name: dbProfil.Name}, nil
}
