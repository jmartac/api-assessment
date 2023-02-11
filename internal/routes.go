package internal

import (
	"api-assessment/internal/controllers"
	"api-assessment/internal/middleware"
	"github.com/gorilla/mux"
)

// NewRouter is used to create and initialise a new router
func NewRouter(services *Services) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LogAPICalls)
	router.Use(middleware.JsonHeader)

	initRoutes(router, services)

	return router
}

// initRoutes is used to initialise the routes, binding them to their respective controllers
func initRoutes(router *mux.Router, services *Services) {
	// Controllers
	filmsController := controllers.NewFilmsController(services.FilmService)
	usersController := controllers.NewUsersController(services.UserService)

	// Films
	router.HandleFunc("/films", filmsController.FindAll).Methods("GET")
	router.HandleFunc("/films/{id}", filmsController.FindByID).Methods("GET")
	router.HandleFunc("/films", filmsController.Create).Methods("POST")
	router.HandleFunc("/films/{id}/update", filmsController.Update).Methods("POST")
	router.HandleFunc("/films/{id}/delete", filmsController.Delete).Methods("POST")

	// Users
	router.HandleFunc("/register", usersController.Create).Methods("POST")
}
