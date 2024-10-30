package handler

import (
	"encoding/json"
	"groupie/fetch"
	"groupie/get"
	"groupie/models"
	"groupie/vars"
	"log"
	"net/http"
	"strconv"
)

// GetArtists fetches all the artists from the api and stores them in an array of objects
func GetArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) > 0 || r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	artists, err := get.GetArtistsData()
	if err != nil {
		vars.Templates.ExecuteTemplate(w, "errors.html", "Unable to fetch artists. Please try again later.")
		return
	}

	if len(artists) == 0 {
		vars.Templates.ExecuteTemplate(w, "errors.html", "No artists found.")
		return
	}

	vars.Templates.ExecuteTemplate(w, "artists.html", artists)
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

	artists, err := get.GetArtistsData()
	if err != nil {
		vars.Templates.ExecuteTemplate(w, "errors.html", "Unable to fetch artists. Please try again later.")
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

	var artistDetails models.ArtistDetails
	artist := artists[artistId-1]
	artistDetails.Artist = artist

	datesBody, datesBody_err := fetch.Fetch(artist.ConcertDates)
	if datesBody_err != nil {
		log.Println("an error occured while fetching artist's dates: ", datesBody_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	datesUnmarshal_err := json.Unmarshal(datesBody, &artistDetails.Dates)
	if datesUnmarshal_err != nil {
		log.Println("an error occured while unmarshalling artist's dates: ", datesUnmarshal_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	relationsBody, relationsBody_err := fetch.Fetch(artist.Relations)
	if relationsBody_err != nil {
		log.Println("an error occured while fetching artist's relations: ", relationsBody_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	relationsUnmarshal_err := json.Unmarshal(relationsBody, &artistDetails.Relations)
	if relationsUnmarshal_err != nil {
		log.Println("an error occured while unmarshalling artist's relations: ", relationsUnmarshal_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	locationsBody, locationsBody_err := fetch.Fetch(artist.Locations)
	if locationsBody_err != nil {
		log.Println("an error occured while fetching artist's locations: ", locationsBody_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	locationsUnmarshal_err := json.Unmarshal(locationsBody, &artistDetails.Locations)
	if locationsUnmarshal_err != nil {
		log.Println("an error occured while unmarshalling artist's locations: ", locationsUnmarshal_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	// fmt.Println(artistDetails)
	vars.Templates.ExecuteTemplate(w, "artistDetails.html", artistDetails)
}
