package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	templates     *template.Template
	artists_url   = "https://groupietrackers.herokuapp.com/api/artists"
	locations_url = "https://groupietrackers.herokuapp.com/api/locations"
)

// struct model for artist's details, fetched using json tags
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type ArtistDetails struct {
	ArtistsName Artist
	Locations   LocationsData
	Dates       ConcertDate
	Relation    Relation
}

type LocationsData struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	//Dates     string   `json:"dates"`
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

var template_dir = "./web/templates/"

func Fetch(url string) ([]byte, error) {
	body := []byte{}
	var body_err error

	response, artists_err := http.Get(url)
	if artists_err != nil {
		return nil, fmt.Errorf("error making a get request to the artists api endpoint: %s", artists_err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, body_err = io.ReadAll(response.Body)
		if body_err != nil {
			return nil, fmt.Errorf("error reading response body: %s", body_err)
		}
	}

	return body, nil
}

// GetArtists fetches all the artists from the api and stores them in an array of objects
func GetArtists(w http.ResponseWriter, r *http.Request) {
	var err error
	templates = template.New("")
	templates, err = templates.ParseGlob(template_dir + "*.html")
	if err != nil {
		log.Fatal(err)
	}
	artists := []Artist{}

	artists_bytes, artists_bytes_err := Fetch(artists_url)
	if artists_bytes_err != nil {
		templates.ExecuteTemplate(w, "errors.html", "Unable to fetch artists. Please try again later.")
		return
	}

	json.Unmarshal(artists_bytes, &artists)

	templates.ExecuteTemplate(w, "artists.html", artists)
}

/*
func GetLocations(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	var err error
	var id int
	id, err = strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	templates = template.New("")
	templates, err = templates.ParseGlob(template_dir + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	locationsResponse := LocationsResponse{
		Index: []LocationsData{},
	}
	fetchedLocations, location_bytes_err := Fetch(locations_url)
	if location_bytes_err != nil {
		templates.ExecuteTemplate(w, "errors.html", "Unable to fetch artist's locations. Please try again later.")
		return
	}

	err = json.Unmarshal(fetchedLocations, &locationsResponse)
	if err != nil {
		log.Fatal(err)
	}
	var DisplayLocations struct {
		Location []string
	}
	for _, location := range locationsResponse.Index {
		if location.Id == id {
			DisplayLocations.Location = location.Locations
			break
		}
	}
	fmt.Println(DisplayLocations)
	templates.ExecuteTemplate(w, "artistDetails.html", DisplayLocations)
}
*/

func MoreDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
		return
	}

	var artists []Artist

	artistsBody, artistsBody_err := Fetch(artists_url)
	if artistsBody_err != nil {
		log.Println("an error occured while fetching artists body: ", artistsBody_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	artistUnmarshal_err := json.Unmarshal(artistsBody, &artists)
	if artistUnmarshal_err != nil {
		log.Println("an error occured while unmarshalling artists body: ", artistUnmarshal_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	idString := r.URL.Query().Get("id")
	artistId, artistId_err := strconv.Atoi(idString)
	if artistId_err != nil || idString == "" {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	if artistId < 1 || artistId > len(artists)-1 {
		http.Error(w, "We did not find an artist with that id.", http.StatusNotFound)
		return
	}

	var ArtistDetails ArtistDetails
	artist := artists[artistId-1]
	var relation Relation
	var dates ConcertDate

	datesBody, datesBody_err := Fetch("https://groupietrackers.herokuapp.com/api/dates/" + idString)
	if datesBody_err != nil {
		log.Println("an error occured while fetching artist's dates: ", datesBody_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	datesUnmarshal_err := json.Unmarshal(datesBody, &dates)
	if datesUnmarshal_err != nil {
		log.Println("an error occured while unmarshalling artist's dates: ", datesUnmarshal_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	relationsBody, relationsBody_err := Fetch("https://groupietrackers.herokuapp.com/api/relation/" + idString)
	if relationsBody_err != nil {
		log.Println("an error occured while fetching artist's relations: ", relationsBody_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	relationsUnmarshal_err := json.Unmarshal(relationsBody, &relation)
	if relationsUnmarshal_err != nil {
		log.Println("an error occured while unmarshalling artist's relations: ", relationsUnmarshal_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	var location LocationsData
	locationsBody, locationsBody_err := Fetch("https://groupietrackers.herokuapp.com/api/locations/" + idString)
	if locationsBody_err != nil {
		log.Println("an error occured while fetching artist's locations: ", locationsBody_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	locationsUnmarshal_err := json.Unmarshal(locationsBody, &location)
	if locationsUnmarshal_err != nil {
		log.Println("an error occured while unmarshalling artist's locations: ", locationsUnmarshal_err)
		templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	
	ArtistDetails.ArtistsName = artist
	ArtistDetails.Locations = location
	ArtistDetails.Dates = dates
	ArtistDetails.Relation = relation

	fmt.Println(ArtistDetails)
}
