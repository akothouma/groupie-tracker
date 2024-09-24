package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var (
	templates    *template.Template
	template_dir = "templates/"
	artists_url  = "https://groupietrackers.herokuapp.com/api/artists"
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

func Fetch(url string) []byte {
	body := []byte{}
	var body_err error

	response, artists_err := http.Get(url)
	if artists_err != nil {
		fmt.Println(artists_err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, body_err = io.ReadAll(response.Body)
		if body_err != nil {
			fmt.Println(body_err)
		}
	}

	return body
}

// GetArtists fetches all the artists from the api and stores them in an array of objects
func GetArtists(w http.ResponseWriter, r *http.Request) {
	artists := []Artist{}

	artists_bytes := Fetch(artists_url)

	json.Unmarshal(artists_bytes, &artists)

	templates.ExecuteTemplate(w, "artists.html", artists)
}

func main() {
	templates, _ = templates.ParseGlob(template_dir + "*.html")

	http.HandleFunc("/", GetArtists)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Println("Listening on :8001...")
	http.ListenAndServe(":8001", nil)
}
