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

// GetArtists handles HTTP GET requests to fetch and display a list of artists.
// It validates the request method and URL path, fetches artist data, and renders
// the artists list to an HTML template.
func GetArtists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorhandling.RenderErr(w, http.StatusMethodNotAllowed, "Not Allowed", "The request method is inappropriate for the URL.")
		return
	}

	if !errorhandling.ClientConnected() {
		errorhandling.RenderErr(w, http.StatusServiceUnavailable, "Connection Error", "Check your internet connection and try again.")
		return
	}

	if len(r.URL.Query()) > 0 || r.URL.Path != "/" {
		errorhandling.RenderErr(w, http.StatusNotFound, "Not found", "The page you are looking for seems not to exist.")
		return
	}

	artists, err := get.GetArtistsData()
	if err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}

	if len(artists) == 0 {
		errorhandling.RenderErr(w, http.StatusNotFound, "That's on our side", "Currently no information to display. Please check in soon.")
		return
	}

	artistsTmplErr := vars.Templates.ExecuteTemplate(w, "artists.html", artists)
	if artistsTmplErr != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
	}
}

// MoreDetails serves an artist's details to a template based on the id provided in the url's query
func MoreDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorhandling.RenderErr(w, http.StatusMethodNotAllowed, "Not Allowed", "The request method is inappropriate for the URL.")
		return
	}

	if !errorhandling.ClientConnected() {
		errorhandling.RenderErr(w, http.StatusServiceUnavailable, "Connection Error", "Check your internet connection and try again.")
		return
	}

	if r.URL.Path != "/artist" {
		errorhandling.RenderErr(w, http.StatusNotFound, "Not found", "The page you are looking for seems not to exist.")
		return
	}

	artists, err := get.GetArtistsData()
	if err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}

	idString := r.URL.Query().Get("id")
	artistId, artistId_err := strconv.Atoi(idString)
	if artistId_err != nil || idString == "" {
		errorhandling.RenderErr(w, http.StatusBadRequest, "Bad Request", "Whoa There! That ID is invalid.")
		return
	}
	if artistId < 1 || artistId > len(artists)-1 {
		errorhandling.RenderErr(w, http.StatusNotFound, "Not Found", "We did not find an artist with that Id.")
		return
	}

	var artistDetails models.ArtistDetails
	artist := artists[artistId-1]
	artistDetails.Artist = artist

	datesBody, datesBody_err := fetch.Fetch(artist.ConcertDates)
	if datesBody_err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}
	datesUnmarshal_err := json.Unmarshal(datesBody, &artistDetails.Dates)
	if datesUnmarshal_err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}

	relationsBody, relationsBody_err := fetch.Fetch(artist.Relations)
	if relationsBody_err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}
	relationsUnmarshal_err := json.Unmarshal(relationsBody, &artistDetails.Relations)
	if relationsUnmarshal_err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}

	locationsBody, locationsBody_err := fetch.Fetch(artist.Locations)
	if locationsBody_err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}
	locationsUnmarshal_err := json.Unmarshal(locationsBody, &artistDetails.Locations)
	if locationsUnmarshal_err != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
		return
	}

	artistDetailsTmplErr := vars.Templates.ExecuteTemplate(w, "artistDetails.html", artistDetails)
	if artistDetailsTmplErr != nil {
		errorhandling.RenderErr(w, http.StatusInternalServerError, "Internal Server Error", "Currently unable to display this page. Please try again later.")
	}
}
