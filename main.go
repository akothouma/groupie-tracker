package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie/handler"
	"groupie/vars"
)

func main() {
	var err error
	vars.Templates, err = vars.Templates.ParseGlob(vars.Template_dir + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler.GetArtists)
	http.HandleFunc("/artist", handler.MoreDetails)

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/static" {
			http.NotFound(w, r)
			return
		}

		fs := http.FileServer(http.Dir("./web/static/"))
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	})

	fmt.Println("Listening on http://localhost:8001/")
	http.ListenAndServe(":8001", nil)
}
