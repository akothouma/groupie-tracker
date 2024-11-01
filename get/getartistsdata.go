package get

import (
	"encoding/json"
	"fmt"

	"groupie/fetch"
	"groupie/models"
	"groupie/vars"
)

// GetArtistsData fetches and returns a list of artists from a specified API endpoint.
// It retrieves the data in JSON format and unmarshals it into a slice of Artist structs.
// Returns: []models.Artist: A slice containing all the artists fetched from the API.
func GetArtistsData() ([]models.Artist, error) {
	artists := []models.Artist{}

	artists_bytes, artists_bytes_err := fetch.Fetch(vars.Artists_url)
	if artists_bytes_err != nil {
		return nil, artists_bytes_err
	}

	unmarshal_err := json.Unmarshal(artists_bytes, &artists)
	if unmarshal_err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", unmarshal_err)
	}

	return artists, nil
}
