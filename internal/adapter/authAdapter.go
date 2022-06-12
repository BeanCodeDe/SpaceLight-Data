package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-AuthMiddleware/authAdapter"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/config"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/dataErr"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const (
	loginUserJson  = `{"Id": "%s", "Password":"%s"}`
	createUserJson = `{"Password":"%s"}`
)

type userCreateDTO struct {
	ID uuid.UUID
}

var (
	serviceToken string
)

func LoginService() {
	log.Debug("Login service")
	authLoginUrl := config.AuthLoginUrl

	loginJson := fmt.Sprintf(loginUserJson, config.ServiceId, config.ServicePassword)
	jsonReq := []byte(loginJson)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, authLoginUrl, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalf("Service couldn't create login request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Service couldn't create login request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status from auth service is not ok but: %v", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Body with token coulnd't be read: %v", err)
	}

	serviceToken = string(bodyBytes)
	log.Debug("Service is logged in")
}

func CreateUser(password string) (uuid.UUID, error) {
	log.Debug("Create user")
	authUserUrl := config.AuthUserUrl

	createJson := fmt.Sprintf(createUserJson, password)
	jsonReq := []byte(createJson)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, authUserUrl, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Errorf("Service couldn't create login request: %v", err)
		return uuid.Nil, dataErr.UnknownError
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set(authAdapter.AuthName, serviceToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Errorf("Service couldn't create login request: %v", err)
		return uuid.Nil, dataErr.UnknownError
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Status from auth service is not ok but: %v", resp.StatusCode)
		return uuid.Nil, dataErr.UnknownError
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Body with token coulnd't be read: %v", err)
		return uuid.Nil, dataErr.UnknownError
	}

	userCreateString := string(bodyBytes)

	var user userCreateDTO
	json.Unmarshal([]byte(userCreateString), &user)
	log.Debug("User Created successful")
	return user.ID, nil
}
