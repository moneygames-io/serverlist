package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/games", Games).Methods("GET")
	log.Fatal(http.ListenAndServe(":6000", router))
}

func Games(w http.ResponseWriter, r *http.Request) {

}
