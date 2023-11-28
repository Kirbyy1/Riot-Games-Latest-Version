// constants.go
package riotclientbuild

import (
	"encoding/json"
	"net/http"
)

// Constants represents the JSON structure returned by the API
type Constants struct {
	Versions VersionsData `json:"data"`
}

// VersionsData represents the nested structure for RiotClientBuild
type VersionsData struct {
	RiotClientBuild string `json:"riotClientBuild"`
}

// GetVersions retrieves the RiotClientBuild version from the specified URL
func GetVersions() (string, error) {
	const VERSION_URL = "https://valorant-api.com/v1/version"

	// Make HTTP GET request
	response, err := http.Get(VERSION_URL)
	if err != nil {
		// Return the default version in case of an error
		return "76.0.4.714.2013", nil
	}
	defer response.Body.Close()

	// Decode JSON response
	var constants Constants
	err = json.NewDecoder(response.Body).Decode(&constants)
	if err != nil {
		// Return the default version in case of JSON decoding error
		return "76.0.4.714.2013", nil
	}

	return constants.Versions.RiotClientBuild, nil
}
