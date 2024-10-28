package main

import (
	"fmt"
	fetch "groupie/Backend"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", fetch.FetchArtists)
	http.HandleFunc("/details", fetch.MoreDetails)
	fmt.Println("Server is running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
