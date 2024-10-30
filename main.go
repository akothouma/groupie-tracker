package main

import (
	"fmt"
	"net/http"
	"strings"

	"groupie/handler"
	"groupie/vars"
)

func main() {
	vars.Templates, _ = vars.Templates.ParseGlob(vars.Template_dir + "*.html")

	http.HandleFunc("/", handler.GetArtists)
	http.HandleFunc("/artist", handler.MoreDetails)

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		referer := r.Header.Get("Referer")
		if referer == "" {
			http.NotFound(w, r)
			return
		}

		fs := http.FileServer(http.Dir("./web/static/"))
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	})

	fmt.Println("Listening on :8001...")
	http.ListenAndServe(":8001", nil)
}
