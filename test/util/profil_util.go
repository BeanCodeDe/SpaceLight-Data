package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/api"
	"github.com/BeanCodeDe/authi/pkg/adapter"
	"github.com/google/uuid"
)

const (
	UserPath = "/profiles/%s"
)

func CreateProfile(userId string, profileCreate *api.ProfileCreateDTO, token string) int {
	response := sendCreateProfileRequest(userId, profileCreate, token)
	defer response.Body.Close()
	return response.StatusCode
}

func sendCreateProfileRequest(userId string, profileCreate *api.ProfileCreateDTO, token string) *http.Response {
	client := &http.Client{}
	profileCreateMarshalled, err := json.Marshal(profileCreate)
	if err != nil {
		panic(err)
	}

	path := Url + fmt.Sprintf(UserPath, userId)
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(profileCreateMarshalled))
	if err != nil {
		panic(err)
	}

	req.Header.Set(CorrelationId, uuid.NewString())
	req.Header.Set(ContentTyp, ContentTypValue)
	req.Header.Set(adapter.AuthorizationHeaderName, "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	return resp
}

func GetProfile(userId string, token string) (int, *api.ProfileResponseDTO) {
	response := sendGetProfileRequest(userId, token)
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return response.StatusCode, nil
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not parse response Body: %v", err)
		return -1, nil
	}

	profileResponse := &api.ProfileResponseDTO{}
	json.Unmarshal(bodyBytes, profileResponse)
	return response.StatusCode, profileResponse
}

func sendGetProfileRequest(userId string, token string) *http.Response {
	client := &http.Client{}

	path := Url + fmt.Sprintf(UserPath, userId)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set(CorrelationId, uuid.NewString())
	req.Header.Set(adapter.AuthorizationHeaderName, "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	return resp
}
