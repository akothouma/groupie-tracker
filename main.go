package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/home":
			handler.GetArtists(w, r)
		case "/details":
			handler.GetLocations(w, r)
	
		default:
			http.Error(w, "404 - Not Found", http.StatusNotFound)
		}
	})

	fs := http.FileServer(http.Dir("./web/static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "styles.css")  {
			http.Error(w, "404 - Not Found", http.StatusNotFound)
			return
		}
		fs.ServeHTTP(w, r)
	})))

	// if err!=nil{
	// 	log.Fatal(err)
	// }
	fmt.Println("Listening on :8001...")
	http.ListenAndServe(":8001", nil)
}
