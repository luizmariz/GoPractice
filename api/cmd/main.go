package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luizmariz/go-practice/api/pkg/db"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/api/mandrake").
		HandlerFunc(handleMandrakeSearch)

	router.
		Methods("GET").
		Path("/api/mandrake").
		HandlerFunc(handleSearchesByUser)

	router.
		Methods("GET").
		Path("/api/mandrake/download").
		HandlerFunc(handleCsvDownload)
}

func handleMandrakeSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("start search")
}

func handleSearchesByUser(w http.ResponseWriter, r *http.Request) {
	log.Println("get user searches")
}

func handleCsvDownload(w http.ResponseWriter, r *http.Request) {
	log.Println("download csv")
}

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed")
	}

	_, err = database.Exec("")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	port := ":8080"
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Println("Server starting...")
	log.Println("Listen on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
