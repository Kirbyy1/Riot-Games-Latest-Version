package version

import (
	"encoding/json"
	"net/http"
)

type VersionResponse struct {
	Status json.RawMessage `json:"status"`
	Data   struct {
		Version string `json:"version"`
	} `json:"data"`
}

func GetVersion() (*VersionResponse, error) {
	REQUEST_URL := "https://valorant-api.com/v1/version"
	res, err := http.Get(REQUEST_URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var versionResponse VersionResponse
	if err := json.NewDecoder(res.Body).Decode(&versionResponse); err != nil {
		return nil, err
	}

	return &versionResponse, nil
}
