package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/api/mandrake").
		HandlerFunc(handleMandrakeSearch)

	// router.
	// 	Methods("GET").
	// 	Path("/api/mandrake").
	// 	HandlerFunc()
}

func handleMandrakeSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("You called a thing!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}