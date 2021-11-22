package db

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type UserDB struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedOn time.Time `db:"created_on"`
	LastLogin time.Time `db:"last_login"`
}

func (user *UserDB) Create() error {
	log.Debugf("Create user %s", user.Name)
	creationTime := time.Now()

	user.CreatedOn = creationTime
	user.LastLogin = creationTime

	if _, err := getConnection().Exec(context.Background(), "INSERT INTO spacelight.user(name,password,created_on,last_login) VALUES($1,MD5($2),$3,$4)", user.Name, user.Password, user.CreatedOn, user.LastLogin); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				log.Warnf("User with name %s already exists", user.Name)
				return echo.NewHTTPError(http.StatusConflict)
			}
		}
		log.Errorf("Unknown error when inserting user: %v", err)
		return echo.ErrInternalServerError
	}
	log.Debugf("User %s inserted into database", user.Name)
	return nil
}

func GetUserByName(username string) (*UserDB, error) {
	log.Debugf("Get user %s by name", username)

	var users []*UserDB
	if err := pgxscan.Select(context.Background(), getConnection(), &users, `SELECT * FROM spacelight.user WHERE name = $1`, username); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.NoDataFound:
				log.Warnf("User with name %s not found", username)
				return nil, echo.NewHTTPError(http.StatusNotFound)
			}
		}
		log.Errorf("Unknown error when getting user by name: %v", err)
		return nil, echo.ErrInternalServerError
	}

	if len(users) != 1 {
		log.Errorf("Cant find only one user. Userlist: %v", users)
		return nil, echo.ErrInternalServerError
	}

	log.Debugf("Got user %v", users[0])
	return users[0], nil
}

func CheckPassword(id uuid.UUID, username string, password string) (bool, error) {
	log.Debugf("Check password for user %s", username)

	var passwordMatches bool
	if err := pgxscan.Select(context.Background(), getConnection(), &passwordMatches,
		`SELECT EXISTS (
		SELECT * FROM spacelight.user WHERE id = $1 AND name = $2 AND password = MD5($3)
	  )`, id, username, password); err != nil {
		log.Errorf("Unknown error when checking password for user %s: %v", username, err)
		return false, echo.ErrInternalServerError
	}

	log.Debugf("Password for user %v is correct", username)
	return passwordMatches, nil
}
