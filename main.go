package main

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie/handler"
)

var (
	templates    *template.Template
	template_dir = "web/templates/"
)

func main() {
	var err error
	templates, err = templates.ParseGlob(template_dir + "*.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler.GetArtists)
	http.HandleFunc("/details", handler.GetLocations)

	fs := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// if err!=nil{
	// 	log.Fatal(err)
	// }
	fmt.Println("Listening on :8001...")
	http.ListenAndServe(":8001", nil)
}
