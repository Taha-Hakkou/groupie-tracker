package main

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/style.css", handlers.CssHandler)
	http.HandleFunc("/", handlers.ArtistsHandler)
	http.HandleFunc("/{id}", handlers.ArtistHandler)
	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
