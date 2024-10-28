package handler

import (
	//"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, r.Method, http.StatusNotFound)
		return
	}

	layoutContent, err := os.ReadFile("Frontend/Layout/pageLayout.html")
	// fmt.Println(layoutContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.New("layout").Parse(string(layoutContent))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	// if err != nil {
	// 	fmt.Print(err)
	// 	os.Exit(1)
	// }
	// var artistsdetails fetch.ArtistDetails

	// decoder := json.NewDecoder(response.Body).Decode(&artistsdetails)
	// temp, err := template.ParseFiles("Frontend/templates/index.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// temp.Execute(w, decoder)
	pageContent, err := os.ReadFile("Frontend/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageIntoLayout := fmt.Sprintf(`{{define "content"}} %s{{end}}`, pageContent)
	tmpl, err = template.Must(tmpl.Clone()).Parse(pageIntoLayout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MoreDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/artistdetails" {
		http.Error(w, r.Method, http.StatusNotFound)
		return
	}

	layoutContent, err := os.ReadFile("Layout/pageLayout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.New("layout").Parse(string(layoutContent))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageContent, err := os.ReadFile("templates/moreInfo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageIntoLayout := fmt.Sprintf(`{{define "content"}} %s{{end}}`, pageContent)
	tmpl, err = template.Must(tmpl.Clone()).Parse(pageIntoLayout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
