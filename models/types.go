package models

// Structs to hold the data returned by the API
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type ArtistDetails struct {
	Artist    Artist
	Locations LocationsData
	Dates     ConcertDate
	Relations Relation
}

type LocationsData struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationsResponse struct {
	Index []LocationsData `json:"index"`
}

type ConcertDate struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
