package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"groupie/fetch"
	"groupie/get"
	"groupie/models"
	"groupie/vars"
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
		w.WriteHeader(500)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Unable to fetch artists. Please try again later.")
		return
	}

	if len(artists) == 0 {
		vars.Templates.ExecuteTemplate(w, "errors.html", "No artists found.")
		return
	}

	vars.Templates.ExecuteTemplate(w, "artists.html", artists)
}

// MoreDetails serves an artist's details to a template based on the id provided in the url's query
func MoreDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path !="/artist"{
		http.Error(w,"MethodNotAllowed",http.StatusMethodNotAllowed)
		return
	}

	artists, err := get.GetArtistsData()
	if err != nil {
		w.WriteHeader(500)
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
		w.WriteHeader(500)
		log.Println("an error occured while fetching artist's dates: ", datesBody_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	datesUnmarshal_err := json.Unmarshal(datesBody, &artistDetails.Dates)
	if datesUnmarshal_err != nil {
		w.WriteHeader(500)
		log.Println("an error occured while unmarshalling artist's dates: ", datesUnmarshal_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	relationsBody, relationsBody_err := fetch.Fetch(artist.Relations)
	if relationsBody_err != nil {
		w.WriteHeader(500)
		log.Println("an error occured while fetching artist's relations: ", relationsBody_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	relationsUnmarshal_err := json.Unmarshal(relationsBody, &artistDetails.Relations)
	if relationsUnmarshal_err != nil {
		w.WriteHeader(500)
		log.Println("an error occured while unmarshalling artist's relations: ", relationsUnmarshal_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	locationsBody, locationsBody_err := fetch.Fetch(artist.Locations)
	if locationsBody_err != nil {
		w.WriteHeader(500)
		log.Println("an error occured while fetching artist's locations: ", locationsBody_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}
	locationsUnmarshal_err := json.Unmarshal(locationsBody, &artistDetails.Locations)
	if locationsUnmarshal_err != nil {
		w.WriteHeader(500)
		log.Println("an error occured while unmarshalling artist's locations: ", locationsUnmarshal_err)
		vars.Templates.ExecuteTemplate(w, "errors.html", "Currently unable to display the requested information. Please try again later.")
		return
	}

	vars.Templates.ExecuteTemplate(w, "artistDetails.html", artistDetails)
}
