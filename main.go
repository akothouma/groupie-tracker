package main

import (
	"fmt"
	"net/http"

	"groupie/handler"
	"groupie/vars"
)

func main() {
	vars.Templates, _ = vars.Templates.ParseGlob(vars.Template_dir + "*.html")

	http.HandleFunc("/", handler.GetArtists)
	http.HandleFunc("/artist", handler.MoreDetails)

	fs := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on :8001...")
	http.ListenAndServe(":8001", nil)
}
