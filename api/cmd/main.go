package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-practice/api/internal/core"
	"go-practice/api/pkg/db"
)

func main() {

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed")
	}

	core := core.NewCore(mux.NewRouter().StrictSlash(true), database)
	port := ":8080"

	core.SetupRouter()

	log.Println("Server starting...")
	log.Println("Listen on port", port)
	log.Fatal(http.ListenAndServe(port, core.Router))
}
