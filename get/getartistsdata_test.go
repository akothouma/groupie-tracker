package get

import (
	"encoding/json"
	"testing"

	"groupie/fetch"
	"groupie/vars"
)

func TestGetArtistsData(t *testing.T) {
	artists, err := GetArtistsData()
	if err != nil {
		t.Errorf("GetArtistsData returned unexpected error: %v", err)
	}

	artists_bytes, artists_bytes_err := fetch.Fetch(vars.Artists_url)
	if artists_bytes_err != nil {
		t.Errorf("unexpected error %v:", artists_bytes_err)
	}

	unmarshal_err := json.Unmarshal(artists_bytes, &artists)
	if unmarshal_err != nil {
		t.Errorf("unexpected error: %v", unmarshal_err)
	}

	for _, artist := range artists {
		if artist.Id == 0 {
			t.Error("unexpected artist id 0")
		}
		if artist.Name == "" {
			t.Error("unexpected empty artist name")
		}
	}
}
