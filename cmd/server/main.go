package main

import (
	"api-assessment/internal"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	services, err := internal.NewServices()
	if err != nil {
		panic(err)
	}

	// Close database connection when main() exits
	defer func(db *internal.Services) {
		log.Println("Disconnecting from database...")
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(services)

	// Migrate database on startup
	err = services.AutoMigrate()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	// TODO Controllers

	// TODO Bind routes to controllers

	log.Println("Starting server...")
	_ = http.ListenAndServe(":3000", router) // TODO use an environment variable
}
