package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	errorhandling "groupie/errorHandling"
	"groupie/fetch"
	"groupie/get"
	"groupie/models"
	"groupie/vars"
)

// GetArtists fetches all the artists from the api and stores them in an array of objects
func GetArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorhandling.RenderErr(w,http.StatusMethodNotAllowed,"Method Not Allowed")
		return
	}

	if len(r.URL.Query()) > 0 || r.URL.Path != "/" {
		errorhandling.RenderErr(w,http.StatusNotFound,"Page Not Found")
		return
	}

	artists, err := get.GetArtistsData()
	if err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Unable to fetch artists. Please try again later.")
		return
	}

	if len(artists) == 0 {
		errorhandling.RenderErr(w, http.StatusBadRequest, "No artist found for listing")
		return
	}

	vars.Templates.ExecuteTemplate(w, "artists.html", artists)
}

// MoreDetails serves an artist's details to a template based on the id provided in the url's query
func MoreDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorhandling.RenderErr(w,http.StatusMethodNotAllowed,"Method Not Allowed")
		return
	}

	if r.URL.Path !="/artist"{
		errorhandling.RenderErr(w,http.StatusNotFound,"Page not found")
		return
	}

	artists, err := get.GetArtistsData()
	if err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Unable to fetch artists. Please try again later.")
		return
	}

	idString := r.URL.Query().Get("id")
	artistId, artistId_err := strconv.Atoi(idString)
	if artistId_err != nil || idString == "" {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	if artistId < 1 || artistId > len(artists)-1 {
		errorhandling.RenderErr(w,http.StatusNotFound,"We did not find an artist with that id.")
		return
	}

	var artistDetails models.ArtistDetails
	artist := artists[artistId-1]
	artistDetails.Artist = artist

	datesBody, datesBody_err := fetch.Fetch(artist.ConcertDates)
	if datesBody_err != nil {
		//log.Println("an error occured while fetching artist's dates: ", datesBody_err)
		errorhandling.RenderErr(w,http.StatusInternalServerError,"Currently unable to display the requested information. Please try again later.")
		return
	}
	datesUnmarshal_err := json.Unmarshal(datesBody, &artistDetails.Dates)
	if datesUnmarshal_err != nil{
		errorhandling.RenderErr(w,http.StatusInternalServerError,"Currently unable to display the requested information. Please try again later.")
		return
	}

	relationsBody, relationsBody_err := fetch.Fetch(artist.Relations)
	if relationsBody_err != nil {
		errorhandling.RenderErr(w,http.StatusInternalServerError,"Currently unable to display the requested information. Please try again later.")
		return
	}
	relationsUnmarshal_err := json.Unmarshal(relationsBody, &artistDetails.Relations)
	if relationsUnmarshal_err != nil {
		errorhandling.RenderErr(w,http.StatusInternalServerError,"Currently unable to display the requested information. Please try again later.")
		return
	}

	locationsBody, locationsBody_err := fetch.Fetch(artist.Locations)
	if locationsBody_err != nil {
		errorhandling.RenderErr(w,http.StatusInternalServerError,"Currently unable to display the requested information. Please try again later.")
		return
	}
	locationsUnmarshal_err := json.Unmarshal(locationsBody, &artistDetails.Locations)
	if locationsUnmarshal_err != nil {
		errorhandling.RenderErr(w,http.StatusInternalServerError,"Currently unable to display the requested information. Please try again later.")
		return
	}

	vars.Templates.ExecuteTemplate(w, "artistDetails.html", artistDetails)
}
