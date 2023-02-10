package main

import (
	"api-assessment/internal"
	"api-assessment/internal/controllers"
	"api-assessment/internal/middleware"
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
	router.Use(middleware.LogAPICalls)
	router.Use(middleware.JsonHeader)

	// Controllers
	filmsController := controllers.NewFilmsController(services.FilmService)

	// Routes
	router.HandleFunc("/films", filmsController.FindAll).Methods("GET")
	router.HandleFunc("/films/{id}", filmsController.FindByID).Methods("GET")
	router.HandleFunc("/film", filmsController.Create).Methods("POST")

	// TODO Bind routes to controllers

	log.Println("Starting server...")
	_ = http.ListenAndServe(":3000", router) // TODO use an environment variable
}
