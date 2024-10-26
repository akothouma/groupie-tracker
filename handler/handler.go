package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
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
type locationsData struct {
	Id int    `json:"id"`
	Locations []string `json:"locations"`
	Dates string `json:"dates"`
}
type LocationsResponse struct {
	Index []locationsData `json:"index"`
}
var template_dir = "./web/templates/"

func Fetch(url string) []byte {
	body := []byte{}
	var body_err error

	response, artists_err := http.Get(url)
	if artists_err != nil {
		log.Printf("The following error was encountered while making a get request to the groupie tracker api: %s", artists_err)
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, body_err = io.ReadAll(response.Body)
		if body_err != nil {
			log.Fatal(body_err)
			return nil
		}
	}

	return body
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

	artists_bytes := Fetch(artists_url)

	json.Unmarshal(artists_bytes, &artists)

	templates.ExecuteTemplate(w, "artists.html", artists)
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	var err error
	templates = template.New("")
	templates, err = templates.ParseGlob(template_dir + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	locationsResponse:= LocationsResponse		{
		Index: []locationsData{},
	}
	fetchedLocations := Fetch(locations_url)

	err=json.Unmarshal(fetchedLocations, &locationsResponse)
	if err!= nil {
        log.Fatal(err)
    }
	fmt.Println(locationsResponse.Index)
	templates.ExecuteTemplate(w, "artistDetails.html", locationsResponse.Index)
}
