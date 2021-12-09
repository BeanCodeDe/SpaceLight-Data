package core

import (
	"SpaceLight/internal/db"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type (
	UserCore struct {
		ID        uuid.UUID
		Name      string
		Password  string
		CreatedOn time.Time
		LastLogin time.Time
	}

	UserCreateDTO struct {
		Name     string `json:"Name" validate:"required,alphanum"`
		Password string `json:"Password" validate:"required"`
	}

	UserLoginDTO struct {
		ID       uuid.UUID `json:"ID"`
		Name     string    `json:"Name" validate:"required,alphanum"`
		Password string    `json:"Password" validate:"required"`
	}

	UserResponseDTO struct {
		ID        uuid.UUID
		Name      string
		CreatedOn time.Time
		LastLogin time.Time
	}
)

func (user *UserCreateDTO) Create() (createdUser *UserResponseDTO, err error) {
	if err = user.mapToUserDB().Create(); err != nil {
		return nil, err
	}

	dbUser, err := db.GetUserByName(user.Name)
	if err != nil {
		return nil, err
	}

	return mapToUserResponseDTO(dbUser), nil
}

func (user *UserLoginDTO) Login() error {
	logedIn, err := db.CheckPassword(user.ID, user.Name, user.Password)
	if err != nil {
		log.Warnf("Could not check password, %v", err)
		return err
	}
	if !logedIn {
		log.Debugf("wrong auth data, %v", user)
		return echo.ErrUnauthorized
	}

	return nil
}

func (userCreateDTO *UserCreateDTO) mapToUserDB() *db.UserDB {
	return &db.UserDB{Name: userCreateDTO.Name, Password: userCreateDTO.Password}
}

func mapToUserResponseDTO(userDB *db.UserDB) *UserResponseDTO {
	return &UserResponseDTO{ID: userDB.ID, Name: userDB.Name, CreatedOn: userDB.CreatedOn, LastLogin: userDB.LastLogin}
}
