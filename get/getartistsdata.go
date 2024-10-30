package get

import (
	"encoding/json"
	"fmt"

	"groupie/fetch"
	"groupie/models"
	"groupie/vars"
)

// GetArtistsData returns an array of all the fetched artists from the api
func GetArtistsData() ([]models.Artist, error) {
	artists := []models.Artist{}

	artists_bytes, artists_bytes_err := fetch.Fetch(vars.Artists_url)
	if artists_bytes_err != nil {
		return nil, artists_bytes_err
	}

	unmarshal_err := json.Unmarshal(artists_bytes, &artists)
	if unmarshal_err != nil {
		return nil, fmt.Errorf("Error unmarshaling JSON: %v", unmarshal_err)
	}

	return artists, nil
}
