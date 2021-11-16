package core

import (
	"SpaceLight/internal/db"
	"time"

	"github.com/google/uuid"
)

type (
	UserCore struct {
		ID        uuid.UUID
		Name      string
		Password  string
		CreatedOn time.Time
		LastLogin time.Time
	}

	UserDTO struct {
		ID       uuid.UUID
		Name     string `json:"Name" validate:"required,alphanum"`
		Password string `json:"Password" validate:"required"`
	}

	UserResponseDTO struct {
		ID        uuid.UUID
		Name      string
		CreatedOn time.Time
		LastLogin time.Time
	}
)

func (user *UserDTO) Create() (createdUser *UserResponseDTO, err error) {
	if err = mapToUserDB(user).Create(); err != nil {
		return nil, err
	}

	dbUser, err := db.GetUserByName(user.Name)
	if err != nil {
		return nil, err
	}

	return mapToUserResponseDTO(dbUser), nil
}

func (user *UserDTO) CheckPassword() (bool, error) {
	return db.CheckPassword(user.ID, user.Name, user.Password)
}

func mapToUserDB(userDTO *UserDTO) *db.UserDB {
	return &db.UserDB{ID: userDTO.ID, Name: userDTO.Name}
}

func mapToUserResponseDTO(userDB *db.UserDB) *UserResponseDTO {
	return &UserResponseDTO{ID: userDB.ID, Name: userDB.Name, CreatedOn: userDB.CreatedOn, LastLogin: userDB.LastLogin}
}
