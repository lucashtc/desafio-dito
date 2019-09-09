package main

import (
	"log"
	"net/http"

	"github.com/lucashtc/desafio-dito/autocomplete"
	"github.com/lucashtc/desafio-dito/timeline"
)

func main() {
	server := autocomplete.ApiServer{}
	server.Port = ":8080"

	http.HandleFunc("/autocomplete_api", server.AutocompleteApi)
	http.HandleFunc("/autocomplete", server.Autocomplete)
	http.HandleFunc("/get", server.Get)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./pages"))))
	http.HandleFunc("/timeline", timeline.TimelineHand)

	log.Fatal(http.ListenAndServe(server.Port, nil))

}
