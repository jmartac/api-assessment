package controllers

import (
	"api-assessment/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

type FilmsController struct {
	fs models.FilmService
}

// NewFilmsController is used to create a new Films controller.
func NewFilmsController(fs models.FilmService) *FilmsController {
	return &FilmsController{fs: fs}
}

// FindAll is used to find all films
// GET /films
func (fc *FilmsController) FindAll(w http.ResponseWriter, _ *http.Request) {
	log.Println("GET /films")
	w.Header().Set("Content-Type", "application/json")

	films, err := fc.fs.FindAll()
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(err)
		if json.NewEncoder(w).Encode(err) != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(films)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
