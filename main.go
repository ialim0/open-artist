package main

import (
	"fmt"
	"groupie-tracker/handle"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	fmt.Printf("Server run on: http://localhost:8080/\n")
	http.HandleFunc("/", handle.HandleHomeRequest)
	http.HandleFunc("/location/", handle.HandleLocationRequest)
	http.HandleFunc("/date/", handle.HandleDatesRequest)
	http.HandleFunc("/relation/", handle.HandleRelationRequest)
	http.HandleFunc("/locations/", handle.HandleAllLocationRequest)
	http.HandleFunc("/locations-artist/", handle.HandleArtistLocationRequest)
	http.HandleFunc("/dates/", handle.HandleAllDatesRequest)
	http.HandleFunc("/dates-artist/", handle.HandleArtistDateRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
