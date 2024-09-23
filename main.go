package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"html/template"
)

var (
	templates    *template.Template
	template_dir = "templates/"
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

// GetArtists fetches all the artists from the api and stores them in an array of objects
func GetArtists(w http.ResponseWriter, r *http.Request) {
	artists_url := "https://groupietrackers.herokuapp.com/api/artists"

	response, artists_err := http.Get(artists_url)
	if artists_err != nil {
		fmt.Println(artists_err)
	}

	defer response.Body.Close()

	artists := []Artist{}

	if response.StatusCode == http.StatusOK {
		artists_bytes, artists_bytes_err := io.ReadAll(response.Body)
		if artists_bytes_err != nil {
			fmt.Println(artists_bytes_err)
		}

		json.Unmarshal(artists_bytes, &artists)
	}

	templates.ExecuteTemplate(w, "artists.html", artists)
}

func main() {
	templates, _ = templates.ParseGlob(template_dir + "*.html")

	http.HandleFunc("/artists", GetArtists)

	fmt.Println("Listening on :8001...")
	http.ListenAndServe(":8001", nil)
}
