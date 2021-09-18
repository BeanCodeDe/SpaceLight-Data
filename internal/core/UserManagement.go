package core

import (
	"SpaceLight/internal/db"
	"time"

	"github.com/google/uuid"
)

type (
	UserCore struct {
		ID        uuid.UUID
		UserName  string
		CreatedOn time.Time
		LastLogin time.Time
	}

	UserDTO struct {
		ID        uuid.UUID
		UserName  string `json:"UserName" validate:"required,alphanum"`
		CreatedOn time.Time
		LastLogin time.Time
	}
)

func (user *UserDTO) Create() (createdUser *UserDTO, err error) {
	if err = mapToUserDB(user).Create(); err != nil {
		return nil, err
	}

	dbUser, err := db.GetUserByName(user.UserName)
	if err != nil {
		return nil, err
	}

	return mapToUserDTO(dbUser), nil
}

func mapToUserDB(userDTO *UserDTO) *db.UserDB {
	return &db.UserDB{ID: userDTO.ID, UserName: userDTO.UserName}
}

func mapToUserDTO(userDB *db.UserDB) *UserDTO {
	return &UserDTO{ID: userDB.ID, UserName: userDB.UserName, CreatedOn: userDB.CreatedOn, LastLogin: userDB.LastLogin}
}
