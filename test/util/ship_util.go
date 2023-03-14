package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/api"
	"github.com/BeanCodeDe/authi/pkg/adapter"
	"github.com/google/uuid"
)

const (
	ShipPath = "/ships"
)

func GetShips(token string) (int, []api.ShipTypeDTO) {
	response := sendGetShipsRequest(token)
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return response.StatusCode, nil
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not parse response Body: %v", err)
		return -1, nil
	}

	shipsResponse := []api.ShipTypeDTO{}
	json.Unmarshal(bodyBytes, shipsResponse)
	return response.StatusCode, shipsResponse
}

func sendGetShipsRequest(token string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, Url+ShipPath, nil)
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
